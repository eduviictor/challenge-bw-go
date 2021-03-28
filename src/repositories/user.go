package repositories

import (
	"github.com/challenge-bw-go/src/database"
	"github.com/challenge-bw-go/src/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(user *models.User) error {
	if err := database.Db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func GetUsers() []models.User {
	users := []models.User{}
	database.Db.Find(&users)
	return users
}

func VerifyIfUserExists(user *models.User, c *gin.Context) error{
	if err := database.Db.Where("id = ?", c.Param("id")).First(&user).Error;  err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *models.User, input *models.User) error{
	if err := database.Db.Model(&user).Updates(input).Error; err != nil {
		return err
	}
	return nil
}