package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SubjectRepository interface {
	SearchSubject(page int, name string, semester int64) (*PaginationMetadata, *[]Subject, error)
	getSubjectById(id uint) (*Subject, error)
	CreateSubject(subject *Subject) error
	UpdateSubject(update *Subject) (*Subject, error)
	DeleteSubjectById(id uint) error
}

var repo SubjectRepository

func paginateSubject(c *gin.Context) {
	pageNumber, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	req := PaginateSubjectRequest{}
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	metadata, subjects, err := repo.SearchSubject(pageNumber, req.Name, req.Semester)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
	}

	resp := PaginateSubjectResponse{
		Page:       metadata.Page,
		PerPage:    metadata.PerPage,
		PageCount:  metadata.PageCount,
		TotalCount: metadata.TotalCount,
		Subjects:   make([]PaginateSubjectResponseItem, len(*subjects)),
	}

	for i, subject := range *subjects {
		resp.Subjects[i] = PaginateSubjectResponseItem{
			Id:       subject.Id,
			Name:     subject.Name,
			Semester: subject.Semester,
		}
	}

	c.JSON(http.StatusOK, resp)
}

func getSubjectById(c *gin.Context) {
	subjectId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	subject, err := repo.getSubjectById(uint(subjectId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	c.JSON(http.StatusOK, subject)
}

func createSubject(c *gin.Context) {
	req := CreateSubjectRequest{}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := json.Unmarshal(body, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid requeust"})
		return
	}

	subject := Subject{
		Name:       req.Name,
		Semester:   req.Semester,
		Detail:     req.Detail,
		Instructor: req.Instructor,
	}

	err = repo.CreateSubject(&subject)
	if err != nil {
		fmt.Println(err)
		if err == ErrConflict {
			c.JSON(http.StatusConflict, gin.H{"error": "object contain duplicated key"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		return
	}

	c.JSON(http.StatusCreated, &subject)
}

func updateSubject(c *gin.Context) {
	req := UpdateSubjectRequest{}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := json.Unmarshal(body, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid requeust"})
		return
	}

	subject := Subject{
		Id:         req.Id,
		Name:       req.Name,
		Semester:   req.Semester,
		Detail:     req.Detail,
		Instructor: req.Instructor,
	}

	updatedSubject, err := repo.UpdateSubject(&subject)
	if err != nil {
		if err == ErrEntityNotFound {
			c.JSON(http.StatusBadRequest, gin.H{"error": "record with given ID not found"})
		} else if err == ErrConflict {
			c.JSON(http.StatusConflict, gin.H{"error": "object contain duplicated key"})
		} else {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		return
	}

	c.JSON(http.StatusCreated, updatedSubject)
}

func deleteSubject(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	if err := repo.DeleteSubjectById(uint(id)); err != nil {
		if err == ErrEntityNotFound {
			c.JSON(http.StatusBadRequest, gin.H{"error": "record with given ID not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}
