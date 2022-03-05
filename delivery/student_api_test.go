package delivery

import (
	"bytes"
	"clean-code/model"
	"clean-code/usecase"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
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

type MockResponse struct {
	StatusCode int
	Message    string
	Data       model.Student
}

type MockErrorResponse struct {
	Message string
}

type studentUseCaseMock struct {
	mock.Mock
}

type StudentApiTestSuite struct {
	suite.Suite
	useCaseTest     usecase.IStudentUseCase
	routerTest      *gin.Engine
	routerGroupTest *gin.RouterGroup
}

func (s *studentUseCaseMock) NewRegistration(student model.Student) (*model.Student, error) {
	args := s.Called(student)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Student), args.Error(1)
}

func (s *studentUseCaseMock) FindStudentInfoById(idCard string) (*model.Student, error) {
	args := s.Called(idCard)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Student), args.Error(1)
}

func (suite *StudentApiTestSuite) SetupTest() {
	suite.useCaseTest = new(studentUseCaseMock)
	suite.routerTest = gin.Default()
	suite.routerGroupTest = suite.routerTest.Group("/apimock")
}

func (suite *StudentApiTestSuite) TestStuidentApi_CreateStudent_Success() {
	dummyStudent := dummyStudents[1]
	suite.useCaseTest.(*studentUseCaseMock).On("NewRegistration", dummyStudent).Return(&dummyStudent, nil)
	delivery := NewStudentApi(suite.useCaseTest)
	// Mau ngetest handler, router test group terserah (yang penting atas bawah sama), yang penting handlernya yang ditest
	handler := delivery.(*StudentApi).createStudent
	registration := "/apimock/murid"
	// registration := "/apimock/student"
	studentRoute := suite.routerGroupTest.Group("/murid")
	// studentRoute := suite.routerGroupTest.Group("/student")
	studentRoute.POST("", handler)

	rr := httptest.NewRecorder()
	reqBody, _ := json.Marshal(dummyStudent)
	request, _ := http.NewRequest(http.MethodPost, registration, bytes.NewBuffer(reqBody))
	request.Header.Set("Content-Type", "application/json")

	suite.routerTest.ServeHTTP(rr, request)
	assert.Equal(suite.T(), rr.Code, 201)

	a := rr.Body.String()
	actualStudent := new(MockResponse)
	json.Unmarshal([]byte(a), actualStudent)
	log.Println("after parse", actualStudent)
	assert.Equal(suite.T(), dummyStudent.Name, actualStudent.Data.Name)
}

func (suite *StudentApiTestSuite) TestStuidentApi_GetById_Success() {
	dummyStudent := dummyStudents[0]
	suite.useCaseTest.(*studentUseCaseMock).On("FindStudentInfoById", "1").Return(&dummyStudent, nil)
	delivery := NewStudentApi(suite.useCaseTest)
	handler := delivery.(*StudentApi).getStudentById
	findById := "/apimock/murid/1"
	studentRoute := suite.routerGroupTest.Group("/murid")
	studentRoute.GET("/:idcard", handler)

	rr := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, findById, nil)
	suite.routerTest.ServeHTTP(rr, request)
	assert.Equal(suite.T(), rr.Code, 200)

	a := rr.Body.String()
	actualStudent := new(MockResponse)
	json.Unmarshal([]byte(a), actualStudent)
	assert.Equal(suite.T(), dummyStudent.Name, actualStudent.Data.Name)
}

func TestStudentApiTestSuite(t *testing.T) {
	suite.Run(t, new(StudentApiTestSuite))
}
