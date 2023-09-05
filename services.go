package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SubjectRepository interface {
	GetSubject(query *Subject) error
	CreateSubject(subject *Subject) error
	UpdateSubject(Subject *Subject) error
	DeleteSubject(subject *Subject) error
}

var repo SubjectRepository

func getSubjectById(c *gin.Context) {
	subjectId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	subject := Subject{Id: uint(subjectId)}
	if err := repo.GetSubject(&subject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
	}

	c.JSON(http.StatusOK, subject)
}

func createSubject(c *gin.Context) {
	req := CreateSubjectRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}

func updateSubject(c *gin.Context) {
	req := UpdateSubjectRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}

func deleteSubject(c *gin.Context) {
	req := DeleteSubjectRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
