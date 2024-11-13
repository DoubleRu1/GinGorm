package controller

import (
	"GinGormCRUD/global"
	"GinGormCRUD/models"
	"GinGormCRUD/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	//register for our new student
	var student models.Student
	//bind information to student structure in /models/student.go
	if err := c.ShouldBind(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error() + "\n Can't bind your information!"})
		return
	}
	//encrypt for students password
	password, err := utils.HashPassword(student.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error() + "\n Can't hash password!"})
		return
	}
	student.Password = password
	fmt.Println(student.Name)

	//generate JWT token
	jwtToken, err := utils.GenerateJWT(student.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error() + "\n Can't generate JWT!"})
		return
	}
	//var sameStudent models.Student
	//store in database using gorm
	if err := global.DB.AutoMigrate(&student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error on automigrate"})
		return
	}
	// can't use the same name
	if err := global.DB.Where("name = ?", student.Name).First(&student).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't use the same name!"})
		return
	}
	if err := global.DB.Create(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": " error on DB create"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": jwtToken})

}

func Login(c *gin.Context) {
	var input struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	var student models.Student
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error() + "Can't bind your information!"})
		return
	}
	db := global.DB
	if err := db.Where("name = ?", input.Name).First(&student).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "wrong student"})
		return
	}
	if !utils.CheckPassword(input.Password, student.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "wrong password"})
		return
	}
	jwtToken, err := utils.GenerateJWT(student.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error() + "\n Can't generate JWT!"})
	}
	c.JSON(http.StatusOK, gin.H{"token": jwtToken})
}
