console.log("user.js loaded");
$(document).ready(function () {
    // Bind click events correctly
    $('#unfollow').on('click', function(event) {
        unfollow(event);
    });

    $('#follow').on('click', function(event) {
        follow(event);
    });
    $('#edit-user').on('submit', function(event){
        edit(event)
    });
    $('#update-password').on('submit', function(event){
        updatePassword(event)
    });
    $('#delete-user').on('click', deleteUser);      
});

function unfollow(event) {
    const userID = $(event.target).data('user-id'); // Get user ID from the button
    $(event.target).prop('disabled', true); // Disable the button

    $.ajax({
        url: `/users/${userID}/unfollow`, // Use backticks for template literals
        method: "POST"
    }).done(function() {
        window.location = `/users/${userID}`; // Redirect after successful unfollow
    }).fail(function() {
        Swal.fire("Ops...", "Error on unfollow", "error"); // Show error message
        $(event.target).prop('disabled', false); // Re-enable the button
    });
}

function follow(event) {
    const userID = $(event.target).data('user-id'); // Get user ID from the button
    $(event.target).prop('disabled', true); // Disable the button

    $.ajax({
        url: `/users/${userID}/follow`, // Use backticks for template literals
        method: "POST"
    }).done(function() {
        window.location = `/users/${userID}`; // Redirect after successful follow
    }).fail(function() {
        Swal.fire("Ops...", "Error on follow", "error"); // Show error message
        $(event.target).prop('disabled', false); // Re-enable the button
    });
}
function edit(event) {
    event.preventDefault();

    $.ajax({
        url:"/edit-user",
        method: "PUT",
        data: JSON.stringify({
            name: $('#name').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
        }),
    }).done(function(){
        Swal.fire("Success", "User updated", "success")
            .then(function(){
                window.location="/profile"
            })
    }).fail(function(){
        Swal.fire("Ops", "Error on update", "error")
    })

}
function updatePassword(event) {
    event.preventDefault();

    if($('#new-password').val() != $('#confirm-password').val()){
        Swal.fire("Ops", "Passwords are different","warning");
        return;
    }
    $.ajax({
        url:"/update-password",
        method: "POST",
        data: JSON.stringify({
            current: $('#current-password').val(),
            new: $('#new-password').val(),
        }),
    }).done(function(){
        Swal.fire("Success", "Password updated", "success")
            .then(function(){
                window.location="/profile"
            })
    }).fail(function(){
        Swal.fire("Ops", "Error on update", "error")
    });
}
function deleteUser() {
    console.log("Delete user button clicked");
    Swal.fire({
        title: "Attention!",
        text: "Are you sure you want to disable your account? :(",
        showCancelButton: true,
        confirmButtonText: "Yes, disable it",
        cancelButtonText: "Cancel",
        icon: "warning"
    }).then(function(confirm){
        if (confirm.value){
            $.ajax({
                url:"/delete-user",
                method: "DELETE"
            }).done(function(){
                Swal.fire("Success!", "User disabled successfully", "success")
                .then(function(){
                    window.location = "/logout";
                })
            }).fail(function(){
                Swal.fire("Ops", "An error occurred, could not disble user", "error");
            });
        }
    });
}

