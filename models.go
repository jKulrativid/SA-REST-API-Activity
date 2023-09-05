package main

type Instructor struct {
	FormalNameName string
	Age            int
	Email          string
	OfficeLocation string
}

type Subject struct {
	Id            uint
	Name          string
	Semester      string
	Detail        string
	Prerequisites []Subject
	Instructors   []Instructor
}
