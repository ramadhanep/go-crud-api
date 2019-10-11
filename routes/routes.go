package routes

import (
	controller "crudapi-courses/controller"
	"fmt"
	"log"

	cors "github.com/AdhityaRamadhanus/fasthttpcors"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func index(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome \n")
}
func helloIndex(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome Hello \n")
}
func hello(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome Hello \n")

	param := ctx.UserValue("param")

	if param == "add" {
		fmt.Fprintln(ctx, "param :", param)
	}
}

func Setup() {
	fmt.Println("Routes")

	router := fasthttprouter.New()

	router.GET("/", index)
	router.GET("/hello", helloIndex)
	router.GET("/hello/:param", hello)

	// COURSE CATEGORIES

	router.GET("/courseCategory", controller.IndexCourseCategory)
	router.POST("/saveCourseCategory", controller.SaveCourseCategory)
	// router.PUT("/updateCourseCategory", controller.UpdateCourseCategory)
	// router.DELETE("/deleteCourseCategory/:id", controller.DeleteCourseCategory)

	// COURSE

	router.GET("/course", controller.IndexCourses)
	router.POST("/saveCourse", controller.SaveCourse)
	router.PUT("/updateCourse", controller.UpdateCourse)
	router.DELETE("/deleteCourse/:id", controller.DeleteCourse)

	listenAddr := ":8000"
	fmt.Println("Running App On localhost port", listenAddr)

	withCors := cors.NewCorsHandler(cors.Options{
		// if you leave allowedOrigins empty then fasthttpcors will treat it as "*"
		AllowedOrigins: []string{"http://localhost:8080"}, // Only allow example.com to access the resource
		// if you leave allowedHeaders empty then fasthttpcors will accept any non-simple headers
		AllowedHeaders: []string{"x-something-client", "content-type"}, // only allow x-something-client and Content-Type in actual request
		// if you leave this empty, only simple method will be accepted
		AllowedMethods:   []string{"GET", "POST"}, // only allow get or post to resource
		AllowCredentials: false,                   // resource doesn't support credentials
		AllowMaxAge:      5600,                    // cache the preflight result
		Debug:            true,
	})

	if err := fasthttp.ListenAndServe(listenAddr, withCors.CorsMiddleware(router.Handler)); err != nil {
		log.Fatal("Error in Listen And Serve : %v", err)
	}
}
