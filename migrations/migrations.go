package migrations

import (
	"bitbucket.org/adson14/api_golang_ii.git/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Imovel{}, models.Proprietario{})
	db.Exec("ALTER TABLE proprietarios ALTER COLUMN data_nascimento TYPE date;")

}
