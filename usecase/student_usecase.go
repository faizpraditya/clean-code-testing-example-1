package usecase

import (
	"clean-code/model"
	"clean-code/repository"
)

type IStudentUseCase interface {
	NewRegistration(student model.Student) (*model.Student, error)
	FindStudentInfoById(idCard string) (*model.Student, error)
	// FindAllStudentByPage(limit int, offset int) ([]model.Student, error)
}

type StudentUseCase struct {
	repo repository.IStudentRepository
}

func NewStudentUseCase(studentRepository repository.IStudentRepository) IStudentUseCase {
	return &StudentUseCase{studentRepository}
}

func (s *StudentUseCase) NewRegistration(student model.Student) (*model.Student, error) {
	return s.repo.CreateOne(student)
}

func (s *StudentUseCase) FindStudentInfoById(idCard string) (*model.Student, error) {
	return s.repo.GetOneById(idCard)
}

// func (s *StudentUseCase) FindAllStudentByPage(limit int, offset int) ([]model.Student, error) {
// 	return s.repo.GetAllByPaging(limit, offset)
// }
