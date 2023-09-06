package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := GetDBConnection()
	if err := db.AutoMigrate(&SubjectTable{}); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	repo = NewSubjectRepo(db)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/subjects/search/:page", paginateSubject)
	r.GET("/subjects/:id", getSubjectById)
	r.POST("/subjects", createSubject)
	r.PUT("/subjects", updateSubject)
	r.DELETE("/subjects/:id", deleteSubject)
	r.Run()
}
