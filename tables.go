package main

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type InstructorTable struct {
	gorm.Model
	FormalName     string `gorm:"index"`
	Age            int
	Email          string `gorm:"index,unique"`
	OfficeLocation string
	TaughtSubjects []SubjectTable `gorm:"many2many:subject_instructors;"`
}

type SubjectTable struct {
	gorm.Model
	Name          string `gorm:"index:name_semester_constraint,unique"`
	Semester      int    `gorm:"index:name_semester_constraint,unique"`
	Detail        string `gorm:"type:text"`
	Prerequisites pq.Int64Array
	Instructors   []InstructorTable `gorm:"many2many:subject_instructors;"`
}
