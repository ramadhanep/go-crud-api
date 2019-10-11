package repositories

import (
	"crudapi-courses/models"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

// RETURN 2 PARAM
// JIKAERROR, VARIABLE ERROR ADA ISINYA

// SELAIN ITU DATA DI STRUCT COURSE YANG ADA ISINYA
func FindAllCourse(db *gorm.DB) ([]models.Course, error) {
	courses := []models.Course{}

	// GORM SELECT ALL
	if err := db.Preload("CourseCategory").Find(&courses).Error; err != nil {
		return nil, err
	}

	return courses, nil
}

func SaveCourse(db *gorm.DB, course *models.Course) (uint, error) {
	err := db.Create(&course).Error

	if err != nil {
		log.Println("Error Saving Courses : %v", err)
		return 0, err
	}

	return course.ID, nil
}

func DeleteCourse(db *gorm.DB, course_id int) (uint, error) {
	// FIND BY ID FIRST
	course := models.Course{}

	if db.Find(&course, course_id).RecordNotFound() {
		log.Printf("Data Not Found")
		return 0, nil
	}

	// DELETE
	if err := db.Where("id = ? AND delete_at IS NULL", course_id).Delete(models.Course{}).Error; err != nil {
		log.Printf("Gagal menghapus data user")
	}
	return course.ID, nil
}

func UpdateCourse(db *gorm.DB, course models.Course) (uint, error) {

	// GET DATA FROM DB
	course_db := models.Course{}
	if err := db.Where("id = ?", course.ID).First(&course_db).Error; err != nil {
		fmt.Println("Course Data Not Found")
		return 0, err
	}
	// SET COURSE DB DARI DATA PARAMETER YG DIKIRIM
	course_db.Name = course.Name
	course_db.Description = course.Description
	course_db.PricePerHour = course.PricePerHour
	course_db.Avatar = course.Avatar

	err := db.Save(&course_db).Error

	// CHECK ERROR
	if err != nil {
		fmt.Printf("Failed to updated data %v", err)
		return 0, err
	}

	return course.ID, nil
}
