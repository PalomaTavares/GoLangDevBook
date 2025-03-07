insert into users (name, nick, email, senha)
values
("José", "zezis", "juserobertu@gmail.com", "$2a$10$UWUCQJpZlUocL5Cn8BMlS.WEM6.SGX68JXsvxFTn6oknYOolMaNmG"),
("Thiago", "thi", "thiago@gmail.com", "$2a$10$UWUCQJpZlUocL5Cn8BMlS.WEM6.SGX68JXsvxFTn6oknYOolMaNmG"),
("Samuel", "Samu", "samuel@gmail.com", "$2a$10$UWUCQJpZlUocL5Cn8BMlS.WEM6.SGX68JXsvxFTn6oknYOolMaNmG");

insert into followers(user_id, follower_id)
values
(1,2),
(2,1),
(1,3);

INSERT INTO posts (title, content, author_id)
VALUES
("Bate o tambô", "meu pai vem de aruanda e a nossa mãe é iansã", 1),
("O Bloco", "meu coração viajou", 2),
("hard truths", "4 estrelas e meia", 3);