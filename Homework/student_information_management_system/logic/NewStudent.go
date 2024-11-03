package logic

import "stuSys/models"

func NewStudent(id int, name, class string) *models.Student {
	return &models.Student{
		Id:    id,
		Name:  name,
		Class: class,
	}
}
