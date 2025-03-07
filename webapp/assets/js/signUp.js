$(document).ready(function() {
    $('#new-account').on('submit', createUser);
});


function createUser(event) {
    event.preventDefault();
    if ($('#password').val() != $('#confirm-password').val()) {
        Swal.fire("Ops", "the passwords are different", "error");
        return;
    }

    $.ajax({
        url:"/users",
        method:"POST",
        contentType: "application/json",
        data: JSON.stringify({
            name: $('#name').val(), 
            email: $('#email').val(),
            nick: $('#nick').val(),
            senha: $('#password').val()
        })
    }).done(function(){
        Swal.fire("Success", "SignUp successful","success" ).then(function(){
            $.ajax({
                url:"/login",
                method: "POST",
                contentType: "application/json",
                data: JSON.stringify({
                    email: $('#email').val(),
                    senha: $('#password').val()
                })
            }).done(function(){
                window.location="/home"
            }).fail(function(error){
                Swal.fire("Ops", "Erron on signUp","error" )
            })
        })

    }).fail(function(error){
        Swal.fire("Ops", "Erron on signUp","error" )
    });
}