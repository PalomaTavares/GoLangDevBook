console.log("✅ login.js carregado!");

$(document).ready(function() {
    $('#login').on('submit', login);
});


function login(event) {
    event.preventDefault();
    console.log("✅ preventDefault chamado!");


    $.ajax({
        url:"/login",
        method:"POST",
        contentType: "application/json",
        data: JSON.stringify({
            email: $('#email').val(),
            senha: $('#senha').val()
        })
    }).done(function(response){
        window.location.href = "/home";
    })
    .fail(function(error){
        Swal.fire("Ops", "Login failed :( invalid user or password","error" )
    });
    
}