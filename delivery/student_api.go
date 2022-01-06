package delivery

import (
	"clean-code/model"
	"clean-code/usecase"
	"clean-code/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StudentApi struct {
	usecase usecase.IStudentUseCase
}

func NewStudentApi(usecase usecase.IStudentUseCase) IDelivery {
	studentApi := StudentApi{
		usecase: usecase,
	}
	return &studentApi
}

func (api *StudentApi) InitRouter(publicRoute *gin.RouterGroup) {
	userRoute := publicRoute.Group("/student")
	// userRoute.GET("/:idcard", utils.GetToken, api.getStudentById)
	userRoute.POST("", api.createStudent)
}

func (api *StudentApi) getStudentById(c *gin.Context) {
	name := c.Param("idcard")
	token := c.Request.Header["Token"]
	fmt.Println("Token nilainya = ", token)
	student, err := api.usecase.FindStudentInfoById(name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": student,
	})
}

// func (api *StudentApi) createStudent(c *gin.Context) {
func (s *StudentApi) createStudent(c *gin.Context) {
	var studentJSON model.StudentJSON
	var student model.Student
	err := c.ShouldBindJSON(&studentJSON)

	if err != nil {
		c.JSON(500, gin.H{"statusCode": 500, "message": "Internal server error"})
		return
	}

	joinDate := utils.StringtoDate(studentJSON.JoinDate, "2006-01-02")
	student.Name = studentJSON.Name
	student.Age = studentJSON.Age
	student.JoinDate = joinDate
	student.Gender = studentJSON.Gender
	student.IdCard = studentJSON.IdCard
	student.Senior = studentJSON.Senior

	// student.JoinDate = time.Now()
	registered, _ := s.usecase.NewRegistration(student)
	// fmt.Println(student)
	// fmt.Println(registered)
	c.JSON(201, gin.H{"statusCode": 201, "message": "Created", "data": registered})
}
