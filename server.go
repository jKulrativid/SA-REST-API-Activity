package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/subjects/:id", getSubjectById)
	r.POST("/subjects", createSubject)
	r.PUT("/subjects", updateSubject)
	r.DELETE("/subjects", deleteSubject)
	r.Run()
}
