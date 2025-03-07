package repositories

import (
	"api/src/models"
	"database/sql"
)

type Posts struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *Posts {
	return &Posts{db}
}

// inserts post on database
func (repository Posts) Create(post models.Post) (uint64, error) {
	statement, error := repository.db.Prepare("insert into posts (title, content, author_id) values (?, ?, ?)")

	if error != nil {
		return 0, error
	}
	defer statement.Close()

	result, error := statement.Exec(post.Title, post.Content, post.IDAuthor)

	if error != nil {
		return 0, error
	}

	lastInsertedID, error := result.LastInsertId()

	if error != nil {
		return 0, error
	}

	return uint64(lastInsertedID), nil

}
func (repository Posts) GetByID(IDPost uint64) (models.Post, error) {
	line, error := repository.db.Query(`
		select p.*, u.nick from
		posts p inner join users u
		on u.id = p.author_id where p.id = ?
	`, IDPost)
	if error != nil {
		return models.Post{}, error
	}

	defer line.Close()

	var post models.Post

	if line.Next() {
		if error = line.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.IDAuthor,
			&post.Likes,
			&post.CreatedIn,
			&post.NickAuthor,
		); error != nil {
			return models.Post{}, error
		}
	}
	return post, nil
}
func (repository Posts) Get(userID uint64) ([]models.Post, error) {
	lines, error := repository.db.Query(`
		select distinct p.*, u.nick from posts p
		inner join users u on u.id = p.author_id
		inner join followers f on p.author_id = f.user_id
		where u.id = ? or f.follower_id = ?
		order by 1 desc`, userID, userID,
	)
	if error != nil {
		return nil, error
	}
	defer lines.Close()
	var posts []models.Post

	for lines.Next() {
		var post models.Post

		if error = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.IDAuthor,
			&post.Likes,
			&post.CreatedIn,
			&post.NickAuthor,
		); error != nil {
			return nil, error
		}

		posts = append(posts, post)
	}
	return posts, nil
}
func (repository Posts) Update(postID uint64, post models.Post) error {
	statement, error := repository.db.Prepare("update posts set title = ?, content = ? where id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error = statement.Exec(post.Title, post.Content, postID); error != nil {
		return error
	}
	return nil
}

func (repository Posts) Delete(postID uint64) error {
	statement, error := repository.db.Prepare("delete from posts where id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error = statement.Exec(postID); error != nil {
		return error
	}
	return nil

}
func (repository Posts) GetByUserID(userID uint64) ([]models.Post, error) {

	lines, error := repository.db.Query(`
	select p.*, u.nick from posts p
	join users u on u.id = p.author_id
	where p.author_id = ?`,
		userID,
	)
	if error != nil {
		return nil, error
	}

	defer lines.Close()

	var posts []models.Post

	for lines.Next() {
		var post models.Post

		if error = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.IDAuthor,
			&post.Likes,
			&post.CreatedIn,
			&post.NickAuthor,
		); error != nil {
			return nil, error
		}

		posts = append(posts, post)
	}
	return posts, nil

}

func (repository Posts) Like(postID uint64) error {
	statement, error := repository.db.Prepare(`update posts set likes = likes + 1 where id = ?`)
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error = statement.Exec(postID); error != nil {
		return error
	}
	return nil
}

func (repository Posts) Unlike(postID uint64) error {
	statement, error := repository.db.Prepare(`update posts set likes =
		CASE 
			WHEN likes > 0 THEN likes - 1
			ELSE 0
		END
		where id = ?
	`)
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error = statement.Exec(postID); error != nil {
		return error
	}
	return nil
}
