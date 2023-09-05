package main

import "gorm.io/gorm"

type subjectRepo struct {
	db *gorm.DB
}

func NewSubjectRepo(db *gorm.DB) SubjectRepository {
	return &subjectRepo{db: db}
}

func (r *subjectRepo) GetSubject(query *Subject) error {
	return nil
}

func (r *subjectRepo) CreateSubject(subject *Subject) error {
	return nil
}

func (r *subjectRepo) UpdateSubject(subject *Subject) error {
	return nil
}

func (r *subjectRepo) DeleteSubject(subject *Subject) error {
	return nil
}
