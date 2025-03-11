package controllers

import (
	"RESTAPICRUD/config"
	"RESTAPICRUD/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario sin autenticacion"})
		return
	}
	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usario no Encontrado"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}
