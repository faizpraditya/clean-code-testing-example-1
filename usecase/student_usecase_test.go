package usecase

import (
	"clean-code/model"
	"clean-code/repository"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var dummyStudents = []model.Student{
	{
		Id:       1,
		Name:     "Dummy name 1",
		Gender:   "M",
		Age:      1,
		JoinDate: time.Time{},
		IdCard:   "dummy id card 1",
		Senior:   false,
	},
	{
		Id:       2,
		Name:     "Dummy name 2",
		Gender:   "F",
		Age:      2,
		JoinDate: time.Time{},
		IdCard:   "dummy id card 2",
		Senior:   true,
	},
}

type repoMock struct {
	mock.Mock
}

func (r *repoMock) CreateOne(student model.Student) (*model.Student, error) {
	args := r.Called(student)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Student), nil
	// return args.Get(0).(*model.Student), args.Error(1)
}

func (r *repoMock) GetOneById(idCard string) (*model.Student, error) {
	args := r.Called(idCard)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Student), args.Error(1)
}

func (r *repoMock) GetOneByName(name string) ([]model.Student, error) {
	args := r.Called(name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Student), args.Error(1)
}

func (r *repoMock) GetAll() ([]model.Student, error) {
	args := r.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Student), args.Error(1)
}

type StudentUseCaseTestSuite struct {
	suite.Suite
	repoTest repository.IStudentRepository
}

func (suite *StudentUseCaseTestSuite) SetupTest() {
	suite.repoTest = new(repoMock)
}

// Melihat behavior usecasenya, kalau sukses ngapain, kalau error ngapain, ga boleh manggil fungsi lain, fokus fungsi tersebut.
func (suite *StudentUseCaseTestSuite) TestStudentUseCaseTestSuite_NewRegistration_Success() {
	dummyStudent := dummyStudents[0]
	suite.repoTest.(*repoMock).On("CreateOne", dummyStudent).Return(&dummyStudent, nil)
	studentUseCaseTest := NewStudentUseCase(suite.repoTest)
	student, err := studentUseCaseTest.NewRegistration(dummyStudent)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummyStudent.Id, student.Id)
}

func (suite *StudentUseCaseTestSuite) TestStudentUseCaseTestSuite_NewRegistration_Failed() {
	dummyStudent := dummyStudents[0]
	suite.repoTest.(*repoMock).On("CreateOne", dummyStudent).Return(nil, errors.New("failed"))
	studentUseCaseTest := NewStudentUseCase(suite.repoTest)
	student, err := studentUseCaseTest.NewRegistration(dummyStudent)
	// assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "failed", err.Error())
	assert.Nil(suite.T(), student)
}

func (suite *StudentUseCaseTestSuite) TestStudentUseCaseTestSuite_FindById_Success() {
	dummyStudent := dummyStudents[0]
	suite.repoTest.(*repoMock).On("GetOneById", dummyStudent.IdCard).Return(&dummyStudent, nil)
	studentUseCaseTest := NewStudentUseCase(suite.repoTest)
	student, err := studentUseCaseTest.FindStudentInfoById("dummy id card 1")
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummyStudent.Id, student.Id)
}

func (suite *StudentUseCaseTestSuite) TestStudentUseCaseTestSuite_FindById_Failed() {
	dummyStudent := dummyStudents[0]
	suite.repoTest.(*repoMock).On("GetOneById", dummyStudent.IdCard).Return(nil, errors.New("failed"))
	studentUseCaseTest := NewStudentUseCase(suite.repoTest)
	student, err := studentUseCaseTest.FindStudentInfoById(dummyStudent.IdCard)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "failed", err.Error())
	assert.Nil(suite.T(), student)
}

func TestStudentUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(StudentUseCaseTestSuite))
}
