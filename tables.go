package main

import (
	"gorm.io/gorm"
)

type InstructorTable struct {
	gorm.Model
	FormalName     string `gorm:"index"`
	Age            int
	Email          string
	OfficeLocation string
}

type SubjectTable struct {
	gorm.Model
	Name          string `gorm:"index;unique"`
	Semester      int
	Detail        string `type:"text"`
	Prerequisites []SubjectTable
	Instructors   []InstructorTable
}
