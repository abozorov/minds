package handler

import (
	"github.com/gin-gonic/gin"
	"hw21/models"
	"net/http"
	"strconv"
)

var users = []models.User{}
var idCounter = 1

// CreateUser godoc// @Summary Создать пользователя
// @Accept json
// @Produce json
// @Param user body models.User true "User Info"// @Success 201 {object} models.User
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = idCounter
	idCounter++
	users = append(users, user)
	c.JSON(http.StatusCreated, user)
}

// GetUsers godoc// @Summary Получить всех пользователей
// @Produce json
// @Success 200 {array} models.User
// @Router /users [get]
func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

// GetUserByID godoc// @Summary Получить пользователя по ID
// @Produce json
// @Param id path int true "User ID"
// @Failure 404 {object} models.ErrorResponse
// @Success 200 {object} map[string]string
// @Router /users/{id} [get]
func GetUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, u := range users {
		if u.ID == id {
			c.JSON(http.StatusOK, u)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

// UpdateUser godoc// @Summary Обновить пользователя
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.User true "User Info"// @Failure 404 {object} models.ErrorResponse
// @Success 200 {object} map[string]string
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, u := range users {
		if u.ID == id {
			newUser.ID = id
			users[i] = newUser
			c.JSON(http.StatusOK, newUser)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

// DeleteUser godoc// @Summary Удалить пользователя
// @Produce json
// @Param id path int true "User ID"
// @Failure 404 {object} models.ErrorResponse
// @Success 200 {object} map[string]string
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}
