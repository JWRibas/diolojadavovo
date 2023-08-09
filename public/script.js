$(document).ready(function() {
    $('#telefone').mask('(00) 00000-0000');
});


    $("#cadastroForm").submit(function(event) {
    event.preventDefault();

    $.ajax({
    type: "POST",
    url: "/cadastrar",
    data: $(this).serialize(),
    success: function(response) {
    alert("Cadastro realizado com sucesso!");


    $("#nome").val("");
    $("#idade").val("");
    $("#produto").val("");
    $("#telefone").val("");
}
});
});


    function consultarCliente() {
    var nome = $("#consultaNome").val();

    $.ajax({
    type: "GET",
    url: "/consultar",
    data: {nome: nome},
    success: function(response) {
    // Crie um novo elemento <li> com as informações do cliente
    var li = $("<li></li>");
    li.text(response.Nome + " / " + response.Idade + " / - " + response.Tipo + " / " + response.Telefone);


    $("#listaClientes").append(li);
}
});
}
