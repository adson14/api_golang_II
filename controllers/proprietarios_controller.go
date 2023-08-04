package controllers

import (
	"strconv"

	"bitbucket.org/adson14/api_golang_ii.git/database"
	"bitbucket.org/adson14/api_golang_ii.git/models"
	"github.com/gin-gonic/gin"
)

func MostrarProprietarios(c *gin.Context) {

	db := database.GetDatabase()

	var proprietarios []models.Proprietario
	err := db.Find(&proprietarios).Error

	if err != nil {
		c.JSON(400, gin.H{"error": "Nada foi encontrado"})
		return
	}

	c.JSON(200, &proprietarios)

}

func MostrarProprietario(c *gin.Context) {
	id := c.Param("id")

	new_id, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{"error": "Necessário id como inteiro"})
	}

	db := database.GetDatabase()

	var proprietario models.Proprietario

	err = db.First(&proprietario, new_id).Error

	if err != nil {
		c.JSON(400, gin.H{"error": "Nenhum proprietário encontrado " + err.Error()})

		return
	}

	c.JSON(200, &proprietario)

}

func CriarProprietario(c *gin.Context) {
	db := database.GetDatabase()

	var proprietario models.Proprietario

	err := c.ShouldBindJSON(&proprietario)

	if err != nil {
		c.JSON(400, gin.H{"error": "Nenhum dado foi enviado no corpo JSON" + err.Error()})
		return
	}

	err = db.Create(&proprietario).Error

	if err != nil {
		c.JSON(400, gin.H{"error": "Não foi possível salvar o imóvel" + err.Error()})
		return
	}

	c.JSON(200, &proprietario)
}

func AtualizarProprietario(c *gin.Context) {
	db := database.GetDatabase()

	var proprietario models.Proprietario

	err := c.ShouldBindJSON(&proprietario)

	if err != nil {
		c.JSON(400, gin.H{"error": "Nenhum dado foi enviado no corpo JSON" + err.Error()})
		return
	}

	err = db.Save(&proprietario).Error

	if err != nil {
		c.JSON(400, gin.H{"error": "Não foi possível salvar o proprietário" + err.Error()})
		return
	}

	c.JSON(200, &proprietario)
}

func RemoveProprietario(c *gin.Context) {
	id := c.Param("id")

	new_id, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{"error": "Necessário  id como inteiro"})
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Proprietario{}, new_id).Error

	if err != nil {
		c.JSON(400, gin.H{"error": "Não foi possível remover o propeirtário"})
	}

	c.Status(204)
}
