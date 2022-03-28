$('#stop-following').on('click', stopFollowing);
$('#follow').on('click', follow);
$('#edit-user').on('submit', edit);
$('#update-password').on('submit', updatePassword);
$('#delete-user').on('click', deleteUser);

function stopFollowing() {
    const userId = $(this).data('user-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/users/${userId}/stop-following`,
        method: "POST"
    }).done(function() {
        window.location = `/users/${userId}`;
    }).fail(function() {
        Swal.fire("Ops...", "Error when trying to stop following user!", "error");
        $('#stop-following').prop('disabled', false);
    });
}

function follow() {
    const userId = $(this).data('user-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/users/${userId}/follow`,
        method: "POST"
    }).done(function() {
        window.location = `/users/${userId}`;
    }).fail(function() {
        Swal.fire("Ops...", "Error following user!", "error");
        $('#follow').prop('disabled', false);
    });
}

function edit(event) {
    event.preventDefault();

    $.ajax({
        url: "/edit-user",
        method: "PUT",
        data: {
            name: $('#name').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
        }
    }).done(function() {
        Swal.fire("Success!", "User informations updated!", "success")
            .then(function() {
                window.location = "/profile";
            });
    }).fail(function() {
        Swal.fire("Ops...", "Error updating user!", "error");
    });
}

function updatePassword(event) {
    event.preventDefault();

    if ($('#new-password').val() != $('#confirm-password').val()) {
        Swal.fire("Ops...", "Passwords do not match!", "warning");
        return;
    }

    $.ajax({
        url: "/update-password",
        method: "POST",
        data: {
            atual: $('#current-password').val(),
            nova: $('#new-password').val()
        }
    }).done(function() {
        Swal.fire("Sucesso!", "Password has been updated!", "success")
            .then(function() {
                window.location = "/profile";
            })
    }).fail(function() {
        Swal.fire("Ops...", "Error updating password!", "error");
    });
}

function deleteUser() {
    Swal.fire({
        title: "Attention!",
        text: "This action cannot be undone, are you sure?",
        showCancelButton: true,
        cancelButtonText: "Cancel",
        icon: "warning"
    }).then(function(confirm) {
        if (confirm.value) {
            $.ajax({
                url: "/delete-user",
                method: "DELETE"
            }).done(function() {
                Swal.fire("Success!", "User deleted!", "success")
                    .then(function() {
                        window.location = "/logout";
                    })
            }).fail(function() {
                Swal.fire("Ops...", "Error deleting user!", "error");
            });
        }
    })
}