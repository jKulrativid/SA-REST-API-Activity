package main

import (
	"gorm.io/gorm"
)

type SubjectTable struct {
	gorm.Model
	Name       string `gorm:"index:name_semester_constraint,unique"`
	Semester   int64  `gorm:"index:name_semester_constraint,unique"`
	Detail     string `gorm:"type:text"`
	Instructor string `gorm:"varchar(255)"`
}
