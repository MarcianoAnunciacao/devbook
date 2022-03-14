$('#new-publication').on('submit', createPublication);

$(document).on('click', '.like', like);
$(document).on('click', '.deslike', deslike);

$('#update-publication').on('click', updatePublication);
$('.delete-publication').on('click', deletePublication);

function createPublication(event) {
    event.preventDefault();

    $.ajax({
        url: "/publications",
        method: "POST",
        data: {
            titulo: $('#title').val(),
            conteudo: $('#content').val(),
        }
    }).done(function() {
        window.location = "/home";
    }).fail(function() {
        Swal.fire("Ops...", "Error creating publication!", "error");
    })
}

function like(event) {
    event.preventDefault();

    const clickedElement = $(event.target);
    const publicationId = elementoClicado.closest('div').data('publicacao-id');

    clickedElement.prop('disabled', true);
    $.ajax({
        url: `/publications/${publicationId}/like`,
        method: "POST"
    }).done(function() {
        const likesCounter = clickedElement.next('span');
        const likesAmout = parseInt(likesCounter.text());

        likesAmout.text(likesAmout + 1);

        clickedElement.addClass('dislike');
        clickedElement.addClass('text-danger');
        clickedElement.removeClass('like');

    }).fail(function() {
        Swal.fire("Ops...", "Error trying to like publication!", "error");
    }).always(function() {
        elementoClicado.prop('disabled', false);
    });
}

function deslike(event) {
    evento.preventDefault();

    const clickedElement = $(event.target);
    const publicationId = clickedElement.closest('div').data('publicacao-id');

    clickedElement.prop('disabled', true);
    $.ajax({
        url: `/publications/${publicationId}/dislike`,
        method: "POST"
    }).done(function() {
        const likesCounter = clickedElement.next('span');
        const likesAmout = parseInt(likesCounter.text());

        likesCounter.text(likesAmout - 1);

        clickedElement.removeClass('descurtir-publicacao');
        clickedElement.removeClass('text-danger');
        clickedElement.addClass('like-publication');

    }).fail(function() {
        Swal.fire("Ops...", "Error disliking a publication!", "error");
    }).always(function() {
        clickedElement.prop('disabled', false);
    });
}

function updatePublication() {
    $(this).prop('disabled', true);

    const publicationId = $(this).data('publication-id');
    
    $.ajax({
        url: `/publications/${publicationId}`,
        method: "PUT",
        data: {
            titulo: $('#title').val(),
            conteudo: $('#content').val()
        }
    }).done(function() {
        Swal.fire('Succsses!', 'Publications has been updated!', 'success')
            .then(function() {
                window.location = "/home";
            })
    }).fail(function() {
        Swal.fire("Ops...", "Error editing publication!", "error");
    }).always(function() {
        $('#update-publication').prop('disabled', false);
    })
}

function deletePublication(event) {
    event.preventDefault();

    Swal.fire({
        title: "Attention!",
        text: "This action cannot be undone, do you sure?",
        showCancelButton: true,
        cancelButtonText: "Cancel",
        icon: "warning"
    }).then(function(confirm) {
        if (!confirm.value) return;

        const clickedElement = $(event.target);
        const publication = clickedElement.closest('div')
        const publicationId = publication.data('publication-id');
    
        elementoClicado.prop('disabled', true);
    
        $.ajax({
            url: `/publication/${publicationId}`,
            method: "DELETE"
        }).done(function() {
            publication.fadeOut("slow", function() {
                $(this).remove();
            });
        }).fail(function() {
            Swal.fire("Ops...", "Error deleting publication!", "error");
        });
    })

}