package models

import (
	"time"

	"gorm.io/gorm"
)

type Imovel struct {
	ID                 uint           `json:"id" gorm:"primarykey"`
	Titulo             string         `json:"titulo"`
	Descricao          string         `json:"descricao"`
	Tipo               string         `json:"tipo"`
	Endereco           string         `json:"endereco"`
	Cidade             string         `json:"cidade"`
	Bairro             string         `json:"bairro"`
	Valor              string         `json:"valor"`
	Quantidade_visitas string         `json:"quantidade_visitas" gorm:"default:0"`
	Area_total         float64        `json:"area_total"`
	Area_construida    float64        `json:"area_construida"`
	Quartos            int            `json:"quartos"`
	Banheiros          int            `json:"banheiros"`
	Suites             int            `json:"suites"`
	Garagens           int            `json:"garagens"`
	Piscinas           int            `json:"piscinas"`
	Condicao           string         `json:"condicao"`
	Tipo_contrato      string         `json:"tipo_contrato"`
	Comissao           float32        `json:"comissao"`
	ProprietarioID     int            `json:"proprietario_id"`
	Proprietario       Proprietario   `gorm:"foreignKey:ProprietarioID"`
	CreatedAt          time.Time      `json:"created"`
	UpdatedAt          time.Time      `json:"updated"`
	DeletedAt          gorm.DeletedAt `json:"deleted"`
}
