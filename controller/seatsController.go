package controller

import (
	"GinGormCRUD/global"
	"GinGormCRUD/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SeatsList(c *gin.Context) {
	var seats []models.Seat
	db := global.DB
	if err := db.Find(&seats).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No data in database"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": seats})
}

func SetSeat(c *gin.Context) {
	var seat models.Seat
	if err := c.ShouldBindJSON(&seat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := global.DB
	if err := db.AutoMigrate(&models.Seat{}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := db.Create(&seat).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": seat})

}
