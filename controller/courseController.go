package controller

import (
	"crudapi-courses/config"
	"crudapi-courses/helpers"
	"crudapi-courses/models"
	"crudapi-courses/repositories"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/valyala/fasthttp"
)

func IndexCourses(ctx *fasthttp.RequestCtx) {

	data := make(map[string]interface{})

	courses, err := repositories.FindAllCourse(config.DB)

	if err != nil {
		fmt.Println("Error get courses %v", err)
	}

	data["courses"] = courses

	helpers.JSON(ctx, data)
}

func SaveCourse(ctx *fasthttp.RequestCtx) {

	// GET DATA DARI BODY
	postValues := ctx.PostBody()
	fmt.Println(postValues)

	// MEMBUAT VARIABLE MAP
	data := make(map[string]interface{})

	// CONVERT DARI JSON TO STRUCT
	course := models.Course{}
	if err := json.Unmarshal(postValues, &course); err != nil {
		log.Println("Error UnMarshal", err)

		data["message"] = "JSON Field Invalid"
		data["error"] = err.Error()
		helpers.JSON(ctx, data)
	}

	// MEMANGGIL FUNGSI SAVE DI REPO
	_, err := repositories.SaveCourse(config.DB, &course)

	if err != nil {
		data["message"] = "Failed to save course"
		data["error"] = err.Error()
	} else {
		data["message"] = "success"
	}
	helpers.JSON(ctx, data)
}

func DeleteCourse(ctx *fasthttp.RequestCtx) {
	// GET ID
	idValue := fmt.Sprintf("%v", ctx.UserValue("id"))
	// CONVERT TO INTEGER
	course_id, err := strconv.Atoi(idValue)

	if err != nil {
		fmt.Println("Error Convert ID %v", err)
	}

	// MEMANGGIL FUNGSI DELETE DI REPO
	_, err = repositories.DeleteCourse(config.DB, course_id)

	// BUAT VARIABLE MAP SEBAGAI OUTPUT
	data := make(map[string]interface{})

	//CHECK ERROR
	if err != nil {
		data["message"] = "Failed to delete data"
		data["error"] = err.Error()
	} else {
		data["message"] = "success"
	}

	// OUTPUT
	helpers.JSON(ctx, data)
}

func UpdateCourse(ctx *fasthttp.RequestCtx) {
	// GET DATA JSON FROM BODY
	postValues := ctx.PostBody()

	// CREATE MAP VARIABLE
	data := make(map[string]interface{})

	// CONVERT JSON TO STRUCT
	course := models.Course{}
	if err := json.Unmarshal(postValues, &course); err != nil {
		data["message"] = "Invalid JSON"
		helpers.JSON(ctx, data)
	}

	// MEMANGGIL FUNGSI DI REPO
	_, err := repositories.UpdateCourse(config.DB, course)

	// CHECK ERROR
	if err != nil {
		data["message"] = err
	} else {
		data["message"] = "success"
	}

	// SET OUTPUT
	helpers.JSON(ctx, data)
}

// func IndexCourses(ctx *fasthttp.RequestCtx) {

// 	// Membuat MAP

// 	data := make(map[string]interface{})

// 	data["data"] = "Hallo"
// 	data["course"] = []string{"Apayaa", "Pepsodent", "Ciptadent"}

// 	ctx.Response.Header.Set("Content-Type", "application/json")

// 	// Convert Ke JSON
// 	res, err := json.Marshal(data)

// 	if err != nil {
// 		panic(err)
// 	}

// 	// Write Output
// 	ctx.Write(res)

// 	// Set Status HTTP CODE
// 	ctx.SetStatusCode(fasthttp.StatusOK)
// }
