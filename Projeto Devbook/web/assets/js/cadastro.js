const formCriarUsuario = document.getElementById('formulario-cadastro')

formCriarUsuario.onsubmit = CriarUsuario
function CriarUsuario(evento){
    evento.preventDefault()
    console.log("Fui clicado!")

    const senha = document.getElementById('senha')
    const senhaConfirmada = document.getElementById('confirmarSenha')

    if(senha.value != senhaConfirmada.value){
        alert('As senhas não coincidem!')
        return
    }

    $.ajax({
        url:"/usuarios",
        method: "POST",
        data:{
            nome: $('#nome').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
            senha: $('#senha').val()
        }
    }).done(function(){
        alert("Usuário cadastrado com sucesso!")
    }).fail(function(erro){
        console.log(erro)
        alert("Erro ao cadastrar usuário!")
    })
}