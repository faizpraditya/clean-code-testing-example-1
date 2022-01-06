package main

import (
	"clean-code/config"
	"log"
)

func main() {
	// appConfig := config.NewConfig()
	// fmt.Println(appConfig)
	// student := model.Student{
	// 	Name:     "jon",
	// 	Gender:   "pria",
	// 	Age:      22,
	// 	JoinDate: time.Now(),
	// 	IdCard:   "002",
	// 	Senior:   true,
	// }
	// fmt.Println(student)
	// newStudent, err := appConfig.UseCaseManager.StudentUseCase().NewRegistration(student)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(newStudent)

	// findStudentPage, err := appConfig.UseCaseManager.StudentUseCase().FindAllStudentByPage(3, 0)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(findStudentPage)

	// ===============================
	// appConfig := config.NewConfig()
	// defer func() {
	// 	if err := appConfig.InfraManager.SqlDB().Close(); err != nil {
	// 		panic(err)
	// 	}
	// }()
	// routeEngine := appConfig.Routes
	// err := routeEngine.RouterEngine.Run(appConfig.ApiBaseUrl)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// =====
	appConfig := config.NewConfig()
	routerEngine := appConfig.Routes
	err := routerEngine.RouterEngine.Run(appConfig.ApiBaseUrl)
	if err != nil {
		log.Fatal(err)
	}
}
