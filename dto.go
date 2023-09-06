package main

type PaginateSubjectRequest struct {
	Name     string `form:"name"`
	Semester int64  `form:"semester"`
}

type PaginateSubjectResponseItem struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Semester int64  `json:"semester"`
}

type PaginateSubjectResponse struct {
	Page       int                           `json:"page"`
	PerPage    int                           `json:"per_page"`
	PageCount  int64                         `json:"page_count"`
	TotalCount int64                         `json:"total_count"`
	Subjects   []PaginateSubjectResponseItem `json:"subjects"`
}

type CreateSubjectRequest struct {
	Name       string `json:"name" binding:"required"`
	Semester   int64  `json:"semester" binding:"required"`
	Detail     string `json:"detail" binding:"required"`
	Instructor string `json:"instructor" binding:"required"`
}

type UpdateSubjectRequest struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	Semester   int64  `json:"semester"`
	Detail     string `json:"detail"`
	Instructor string `json:"instructor"`
}
