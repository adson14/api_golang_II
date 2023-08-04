package controllers

import (
	"strconv"

	"bitbucket.org/adson14/api_golang_ii.git/database"
	"bitbucket.org/adson14/api_golang_ii.git/models"
	"github.com/gin-gonic/gin"
)

func MostrarImoveis(c *gin.Context) {

	db := database.GetDatabase()

	var imoveis []models.Imovel
	err := db.Find(&imoveis).Error

	if err != nil {
		c.JSON(400, gin.H{"error": "Nada foi encontrado"})
		return
	}

	c.JSON(200, &imoveis)

}

func MostrarImovel(c *gin.Context) {
	id := c.Param("id")

	new_id, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{"error": "Necessário id como inteiro"})
	}

	db := database.GetDatabase()

	var imovel models.Imovel

	err = db.First(&imovel, new_id).Error

	if err != nil {
		c.JSON(400, gin.H{"error": "Nenhum imóvel encontrado " + err.Error()})

		return
	}

	c.JSON(200, &imovel)

}

func CriarImovel(c *gin.Context) {
	db := database.GetDatabase()

	var imovel models.Imovel

	err := c.ShouldBindJSON(&imovel)

	if err != nil {
		c.JSON(400, gin.H{"error": "Nenhum dado foi enviado no corpo JSON" + err.Error()})
		return
	}

	err = db.Create(&imovel).Error

	if err != nil {
		c.JSON(400, gin.H{"error": "Não foi possível salvar o imóvel" + err.Error()})
		return
	}

	c.JSON(200, &imovel)
}

func AtualizarImovel(c *gin.Context) {
	db := database.GetDatabase()

	var imovel models.Imovel

	err := c.ShouldBindJSON(&imovel)

	if err != nil {
		c.JSON(400, gin.H{"error": "Nenhum dado foi enviado no corpo JSON" + err.Error()})
		return
	}

	err = db.Save(&imovel).Error

	if err != nil {
		c.JSON(400, gin.H{"error": "Não foi possível salvar o imóvel" + err.Error()})
		return
	}

	c.JSON(200, &imovel)
}

func RemoveImovel(c *gin.Context) {
	id := c.Param("id")

	new_id, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{"error": "Necessário id como inteiro"})
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Imovel{}, new_id).Error

	if err != nil {
		c.JSON(400, gin.H{"error": "Não foi possível remover o imóvel"})
	}

	c.Status(204)
}
