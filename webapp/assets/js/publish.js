$('#new-publication').on('submit', criarPublicacao);

$(document).on('click', '.like', like);
$(document).on('click', '.deslike', deslike);

$('#update-publication').on('click', updatePublication);
$('.delete-publication').on('click', deletePublication);

function criarPublicacao(event) {
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
        const likesCounter = elementoClicado.next('span');
        const likesAmout = parseInt(likesCounter.text());

        likesAmout.text(likesAmout + 1);

        clickedElement.addClass('dislike');
        clickedElement.addClass('text-danger');
        clickedElement.removeClass('like');

    }).fail(function() {
        Swal.fire("Ops...", "Error when try to like publication!", "error");
    }).always(function() {
        elementoClicado.prop('disabled', false);
    });
}

function deslike(event) {
    evento.preventDefault();

    const elementoClicado = $(evento.target);
    const publicacaoId = elementoClicado.closest('div').data('publicacao-id');

    elementoClicado.prop('disabled', true);
    $.ajax({
        url: `/publicacoes/${publicacaoId}/descurtir`,
        method: "POST"
    }).done(function() {
        const contadorDeCurtidas = elementoClicado.next('span');
        const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());

        contadorDeCurtidas.text(quantidadeDeCurtidas - 1);

        elementoClicado.removeClass('descurtir-publicacao');
        elementoClicado.removeClass('text-danger');
        elementoClicado.addClass('curtir-publicacao');

    }).fail(function() {
        Swal.fire("Ops...", "Erro ao descurtir a publicação!", "error");
    }).always(function() {
        elementoClicado.prop('disabled', false);
    });
}

function atualizarPublicacao() {
    $(this).prop('disabled', true);

    const publicacaoId = $(this).data('publicacao-id');
    
    $.ajax({
        url: `/publicacoes/${publicacaoId}`,
        method: "PUT",
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val()
        }
    }).done(function() {
        Swal.fire('Sucesso!', 'Publicação criada com sucesso!', 'success')
            .then(function() {
                window.location = "/home";
            })
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao editar a publicação!", "error");
    }).always(function() {
        $('#atualizar-publicacao').prop('disabled', false);
    })
}

function deletarPublicacao(evento) {
    evento.preventDefault();

    Swal.fire({
        title: "Atenção!",
        text: "Tem certeza que deseja excluir essa publicação? Essa ação é irreversível!",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then(function(confirmacao) {
        if (!confirmacao.value) return;

        const elementoClicado = $(evento.target);
        const publicacao = elementoClicado.closest('div')
        const publicacaoId = publicacao.data('publicacao-id');
    
        elementoClicado.prop('disabled', true);
    
        $.ajax({
            url: `/publicacoes/${publicacaoId}`,
            method: "DELETE"
        }).done(function() {
            publicacao.fadeOut("slow", function() {
                $(this).remove();
            });
        }).fail(function() {
            Swal.fire("Ops...", "Erro ao excluir a publicação!", "error");
        });
    })

}