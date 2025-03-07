$(document).ready(function () {
    $('#new-post').on('submit', createPost);
    $(document).on('click', '.like-post', likePost)
    $(document).on('click', '.unlike-post', unlikePost)
    $('#update-post').on('click', updatePost)
    $(document).on('click', '.delete-post', deletePost);
});

function createPost(event) {
    event.preventDefault();

    $.ajax({
        url: "/posts",
        method: "POST",
        contentType: "application/json",
        data: JSON.stringify({
            title: $('#title').val(),
            content: $('#content').val(),
        })
    }).done(function () {
        window.location.href = "/home";
    }).fail(function (jqXHR) {
        Swal.fire("Ops", "Erron on create post","error" )
    });
}
function likePost(event) {
    event.preventDefault();
    const clicked = $(event.target);

    // Find the closest div with `data-post-id`
    const postID = clicked.closest('.jumbotron').data('post-id');
    clicked.prop('disabled', true);
    $.ajax({
        url: `/posts/${postID}/like`,
        method: "PUT"
    }).done(function () {
        const countLikes = clicked.next('span');
        const nLikes = parseInt(countLikes.text());
        countLikes.text(nLikes + 1);

        clicked.addClass('unlike-post');
        clicked.addClass('text-danger')
        clicked.removeClass('like-post');

    }).fail(function () {
        Swal.fire("Ops", "Failed to like post","error" )
    }).always(function () {
        clicked.prop('disabled', false);
    });
}
function unlikePost(event) {
    event.preventDefault();
    const clicked = $(event.target);

    console.log('Unlike post clicked!', clicked); // Debugging line

    // Find the closest div with `data-post-id`
    const postID = clicked.closest('.jumbotron').data('post-id');
    clicked.prop('disabled', true);
    $.ajax({
        url: `/posts/${postID}/unlike`,
        method: "PUT"
    }).done(function () {
        const countLikes = clicked.next('span');
        const nLikes = parseInt(countLikes.text());
        countLikes.text(nLikes - 1);

        clicked.removeClass('unlike-post');
        clicked.removeClass('text-danger');
        clicked.addClass('like-post');
    }).fail(function () {
        Swal.fire("Ops", "Failed to unlike post","error" )
    }).always(function () {
        clicked.prop('disabled', false);
    });
}
function updatePost() {
    $(this).prop('disabled', true);
    const postID = $(this).data('post-id');
    console.log(postID);

    $.ajax({
        url: `/posts/${postID}`,
        method: "PUT",
        contentType: "application/json",
        data: JSON.stringify({
            title: $('#title').val(),
            content: $('#content').val()
        }),
        dataType: "json"
    }).done(function () {
        Swal.fire(
            'Success',
            'Post updated!',
            'success'
        ).then(function () {
            window.location = "/home"
        })
    }).fail(function () {
        Swal.fire("Ops", "Failed to update","error" )
    }).always(function () {
        $('#update-post').prop('disabled', false);
    });
}
function deletePost(event) {
    event.preventDefault();
    Swal.fire({
        title: "Attention!",
        text: "Are you sure you want to delete?",
        showCancelButton: true,
        cancelButtonText: "Cancel",
        icon: "warning"
    }).then(function(confirm) {
        if (!confirm.value) return;

        const clicked = $(event.target);
        const post = clicked.closest('.jumbotron')
        const postID = post.data('post-id');

        if (!postID) {
            console.log("⚠️ Post ID not found.");
            return;
        }
        clicked.prop('disabled', true);

        $.ajax({
            url: `/posts/${postID}`,
            method: "DELETE",
        }).done(function () {
            post.fadeOut("slow", function () {
                $(this).remove();
            });
        }).fail(function (jqXHR) {
            Swal.fire("Ops", "Erron on delete","error" )
        });
    })

}