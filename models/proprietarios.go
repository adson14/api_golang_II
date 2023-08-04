package models

import "time"

type Proprietario struct {
	ID              uint      `json:"id" gorm:"primarykey"`
	Nome            string    `json:"nome"`
	CPF             string    `json:"cpf"`
	RG              string    `json:"rg"`
	Data_nascimento time.Time `json:"data_nascimento" gorm:"type:date;Column:data_nascimento"`
	Estado_civil    string    `json:"estado_civil"`
	Telefone        string    `json:"telefone"`
	Email           string    `json:"email"`
	Endereco        string    `json:"endereco"`
	Numero          string    `json:"numero"`
	Bairro          string    `json:"bairro"`
	Cidade          string    `json:"cidade"`
	Estado          string    `json:"estado"`
}
