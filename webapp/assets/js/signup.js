$('#signin-form').on('submit', createUser);

function createUser(event) {
    event.preventDefault();

    if ($('#password').val() != $('#confirm-password').val()) {
        Swal.fire("Ops...", "Please check password!", "error");
        return;
    }

    $.ajax({
        url: "/users",
        method: "POST",
        data: {
           nome: $('#name').val(), 
           email: $('#email').val(),
           nick: $('#nick').val(),
           senha: $('#password').val()
        }
    }).done(function() {
        Swal.fire("Success!", "User created!", "success")
            .then(function() {
                $.ajax({
                    url: "/login",
                    method: "POST",
                    data: {
                        email: $('#email').val(),
                        senha: $('#password').val()
                    }
                }).done(function() {
                    window.location = "/home";
                }).fail(function() {
                    Swal.fire("Ops...", "Authentication failed", "error");
                })
            })
    }).fail(function() {
        Swal.fire("Ops...", "Error creating user!", "error");
    });
}
