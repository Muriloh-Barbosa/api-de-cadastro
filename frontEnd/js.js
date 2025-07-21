function validarCPF(cpf) {
  const regex = /^\d{11}$/;
  return regex.test(cpf);
}

function mostrarMensagem(texto, sucesso = false) {
  const $msg = $("#mensagem");
  $msg
    .css("color", sucesso ? "#28a745" : "#d9534f")
    .text(texto)
    .fadeIn();

  // Limpa mensagem após 10 segundos
  setTimeout(() => {
    $msg.fadeOut(300, () => {
      $msg.text("").show();
    });
  }, 10000);
}

function montarPayload() {
  return {
    nomeCompleto: $("#nomeCompleto").val().trim(),
    cpf: $("#cpf").val().trim(),
    Nascimento: $("#Nascimento").val(),
    Endereco: $("#Endereco").val().trim(),
    Cidade: $("#Cidade").val().trim(),
    Estado: $("#Estado").val().trim(),
    Telefone: $("#Telefone").val().trim(),
    Email: $("#Email").val().trim(),
    Codigo: $("#Codigo").val().trim()
  };
}

function preencherFormulario(data) {
  Object.keys(data).forEach(k => {
    $("#" + k).val(data[k]);
  });
}

// CREATE
$("#button-create").click(() => {
  const dados = montarPayload();
  if (!validarCPF(dados.cpf)) {
    mostrarMensagem("CPF inválido.");
    return;
  }
  $.ajax({
    url: "http://localhost:8080/create/user",
    type: "POST",
    data: JSON.stringify(dados),
    contentType: "application/json",
    success: (res) => {
      mostrarMensagem(res.message || "Cadastro criado com sucesso!", true);
      $("#form-cadastro")[0].reset();
      $("#resultado").empty();
    },
    error: (xhr) => {
      const resposta = xhr.responseJSON;
      mostrarMensagem(resposta?.error || "Erro ao criar cadastro.");
    }
  });
});

// UPDATE
$("#button-update").click(() => {
  const dados = montarPayload();
  if (!validarCPF(dados.cpf)) {
    mostrarMensagem("CPF inválido.");
    return;
  }
  $.ajax({
    url: "http://localhost:8080/update/user",
    type: "PUT",
    data: JSON.stringify(dados),
    contentType: "application/json",
    success: (res) => {
      mostrarMensagem(res.message || "Cadastro atualizado com sucesso!", true);
    },
    error: (xhr) => {
      const resposta = xhr.responseJSON;
      mostrarMensagem(resposta?.error || "Erro ao atualizar cadastro.");
    }
  });
});

// GET por CPF
$("#button-get").click(() => {
  const cpf = $("#cpf-consulta").val().trim();
  if (!cpf) {
    mostrarMensagem("Informe o CPF para consulta.");
    return;
  }
  if (!validarCPF(cpf)) {
    mostrarMensagem("CPF inválido.");
    return;
  }
  $.ajax({
    url: `http://localhost:8080/user?cpf=${cpf}`,
    type: "GET",
    success: (res) => {
      const data = res.data;
      const msg = res.message || "Cadastro encontrado!";

      if (!data || Object.keys(data).length === 0) {
        mostrarMensagem("Cadastro não encontrado.");
        $("#resultado").empty();
        return;
      }

      mostrarMensagem(msg, true);
      preencherFormulario(data);

      let html = "<h3>Dados encontrados:</h3><ul>";
      Object.keys(data).forEach(k => {
        html += `<li><strong>${k}:</strong> ${data[k]}</li>`;
      });
      html += "</ul>";
      $("#resultado").html(html).fadeIn();

      // Esconde os dados após 10 segundos
      setTimeout(() => {
        $("#resultado").fadeOut(300, () => {
          $("#resultado").empty().show();
        });
      }, 10000);
    },
    error: (xhr) => {
      const resposta = xhr.responseJSON;
      mostrarMensagem(resposta?.error || "Erro ao consultar cadastro.");
      $("#resultado").empty();
    }
  });
});
