package structs

type Informacoes struct {
	Nomecompleto string `json:"nomeCompleto" valid:"required"`
	Cpf          string `json:"cpf" valid:"required"`
	Nascimento   string `json:"Nascimento" valid:"optional"`
	Endereco     string `json:"Endereco" valid:"optional"`
	Cidade       string `json:"Cidade" valid:"optional"`
	Estado       string `json:"Estado" valid:"optional"`
	Telefone     string `json:"Telefone" valid:"optional"`
	Email        string `json:"Email" valid:"optional"`
	Codigo       string `json:"Codigo" valid:"required"`
}
