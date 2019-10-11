package main

import (
	"crudapi-courses/config"
	"crudapi-courses/models"
	"crudapi-courses/routes"
)

func main() {

	config.DB.Debug().AutoMigrate(&models.Course{}, &models.CourseCategory{})

	routes.Setup()
}
