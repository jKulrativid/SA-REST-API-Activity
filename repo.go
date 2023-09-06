package main

import (
	"strings"

	"gorm.io/gorm"
)

type subjectRepo struct {
	limit int
	db    *gorm.DB
}

func NewSubjectRepo(db *gorm.DB) SubjectRepository {
	return &subjectRepo{
		db:    db,
		limit: 10,
	}
}

func (r *subjectRepo) SearchSubject(page int, name string, semester int64) (*PaginationMetadata, *[]Subject, error) {
	var subjectTables []SubjectTable
	var totalCount int64

	query := r.db.Model(&SubjectTable{})
	if name != "" {
		query.Where("name = ?", name)
	}

	if semester != 0 {
		query.Where("semester = ?", semester)
	}

	query.Count(&totalCount)

	offset := (page - 1) * r.limit
	result := query.Offset(offset).Limit(r.limit).Find(&subjectTables)
	if err := result.Error; err != nil {
		return nil, nil, err
	}

	metadata := PaginationMetadata{
		Page:       page,
		PerPage:    r.limit,
		PageCount:  result.RowsAffected,
		TotalCount: totalCount,
	}

	subjects := make([]Subject, len(subjectTables))
	for i, subjectTable := range subjectTables {
		subjects[i] = Subject{
			Id:         subjectTable.ID,
			Name:       subjectTable.Name,
			Semester:   subjectTable.Semester,
			Detail:     subjectTable.Detail,
			Instructor: subjectTable.Instructor,
		}
	}
	return &metadata, &subjects, nil
}

func (r *subjectRepo) getSubjectById(id uint) (*Subject, error) {
	subjectTable := SubjectTable{}

	result := r.db.First(&subjectTable, id)
	if err := result.Error; err != nil {
		return nil, err
	}

	subject := Subject{
		Id:         subjectTable.ID,
		Name:       subjectTable.Name,
		Semester:   subjectTable.Semester,
		Detail:     subjectTable.Detail,
		Instructor: subjectTable.Instructor,
	}
	return &subject, nil
}

func (r *subjectRepo) CreateSubject(subject *Subject) error {
	subjectTable := SubjectTable{
		Name:       subject.Name,
		Semester:   subject.Semester,
		Detail:     subject.Detail,
		Instructor: subject.Instructor,
	}

	result := r.db.Create(&subjectTable)
	if err := result.Error; err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return ErrConflict
		}
		return err
	}

	subject.Id = subjectTable.ID

	return nil
}

func (r *subjectRepo) UpdateSubject(subject *Subject) (*Subject, error) {
	subjectTable := SubjectTable{
		Name:       subject.Name,
		Semester:   subject.Semester,
		Detail:     subject.Detail,
		Instructor: subject.Instructor,
	}

	result := r.db.Model(&SubjectTable{}).Where("id = ?", subject.Id).Updates(&subjectTable)
	if result.RowsAffected == 0 {
		return nil, ErrEntityNotFound
	}
	if err := result.Error; err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return nil, ErrConflict
		}
		return nil, err
	}

	updatedSubjectTable := SubjectTable{}
	if err := r.db.First(&updatedSubjectTable, subject.Id).Error; err != nil {
		return nil, err
	}

	return &Subject{
		Id:         updatedSubjectTable.ID,
		Name:       updatedSubjectTable.Name,
		Semester:   updatedSubjectTable.Semester,
		Detail:     updatedSubjectTable.Detail,
		Instructor: updatedSubjectTable.Instructor,
	}, nil
}

func (r *subjectRepo) DeleteSubjectById(id uint) error {
	result := r.db.Unscoped().Delete(&SubjectTable{}, id)
	if result.RowsAffected == 0 {
		return ErrEntityNotFound
	}

	if err := result.Error; err != nil {
		return err
	}

	return nil
}
