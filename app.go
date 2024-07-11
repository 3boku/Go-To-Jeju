package main

import (
	"Go-To-Jeju/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type msgStruct struct {
	Message string `json:"message"`
}

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type"},
		MaxAge:       24 * time.Hour,
	}))

	r.POST("/chat", func(c *gin.Context) {
		var msg msgStruct
		err := c.ShouldBindJSON(&msg)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp := services.ChatWithKikuri(msg.Message)
		c.JSON(http.StatusOK, gin.H{
			"message": resp,
		})
	})

	r.Run(":8080")
}
