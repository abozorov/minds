package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type people struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	serv := gin.Default()

	// 	1. Создать GET эндпоинт /hello, который возвращает JSON: { "message":
	// "Hello, <name>" } с query параметром name.
	serv.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Hello, %s", c.Query("name")),
		})
	})

	// 2. Создать POST эндпоинт /user, принимающий JSON с именем и возрастом, и
	// возвращающий его обратно.
	serv.POST("/user", func(c *gin.Context) {
		var person people
		if err := c.BindJSON(&person); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "error during binding json, error is" + err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, person)
	})

	// 3. Создать path-параметр /user/:id, который возвращает ID пользователя.
	serv.GET("/user/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"id": c.Param("id"),
		})
	})

	err := serv.Run(":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
}
