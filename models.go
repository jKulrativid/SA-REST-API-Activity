package main

type Subject struct {
	Id         uint
	Name       string
	Semester   int64
	Detail     string
	Instructor string
}

type PaginationMetadata struct {
	Page       int
	PerPage    int
	PageCount  int64
	TotalCount int64
}
