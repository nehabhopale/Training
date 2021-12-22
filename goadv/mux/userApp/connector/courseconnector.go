package connector

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"pass/model"
	"pass/services"
	"pass/handler"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)
type courseConnector struct{
	handler  *handler.Handler
	userService     *services.UserService
	courseService *services.CourseService
}
func NewCourseConnector(handler *handler.Handler,userService *services.UserService, courseService *services.CourseService) *courseConnector {
	return &courseConnector{
		handler:handler,
		userService:     userService,
		courseService: courseService,
	}
}

func (c *courseConnector)RegisterCourseRoutes(authRoute *mux.Router,noAuthRoute *mux.Router) {
	noAuthRoute.HandleFunc("/course", c.AddCourse).Methods("POST")
	authRoute.Use(c.handler.ValidAuth)
	authRoute.HandleFunc("/course", c.GetAllCourses).Methods("GET")
	authRoute.HandleFunc("/course",c. GetAllCourses).Queries("limit", "{limit:[0-9]+}", "pageNo", "{pageNo:[0-9]+}").Methods("GET")
	//authRoute.HandleFunc("/course", c.AddCourse).Methods("POST")
	authRoute.HandleFunc("/course/{id}", c.UpdateCourse).Methods("PUT")
	authRoute.HandleFunc("/course/{id}", c.GetCourseFromId).Methods("GET")
	authRoute.HandleFunc("/course/{id}", c.DeleteCourse).Methods("DELETE")
}
func(c *courseConnector) AddCourse(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("User-Count", strconv.Itoa(c.userService.GetUsersCount()))
	var course model.Course
	json.NewDecoder(r.Body).Decode(&course)
	c.courseService.AddCourse(&course)
	// err2:=db.Debug().Model(course).Association("Users").Error
	// if err2!=nil{
	// 	fmt.Println("error in association------>",err2)
	// }
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(course)
	
}

func (c *courseConnector)GetAllCourses(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("User-Count", strconv.Itoa(c.userService.GetUsersCount()))
	limit, _ := strconv.Atoi(r.FormValue("limit"))
	pageNo, _ := strconv.Atoi(r.FormValue("pageNo"))
	offset := limit * (pageNo - 1)
	fmt.Println(limit, pageNo)
	var courses []model.Course
	c.courseService.GetAllCourses(&courses, limit, offset)
	json.NewEncoder(w).Encode(courses)
	
}
func (c *courseConnector)GetCourseFromId(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("User-Count", strconv.Itoa(c.userService.GetUsersCount()))
		values := mux.Vars(r)
		id, _ := uuid.FromString(values["id"])
		var course model.Course
		course.ID=id
		var str1 []string
		c.courseService.GetCourseFromId(&course,id,str1)
		json.NewEncoder(w).Encode(course)

}

func (c *courseConnector)UpdateCourse(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("User-Count", strconv.Itoa(c.userService.GetUsersCount()))
		values := mux.Vars(r)
		id, _ := uuid.FromString(values["id"])
		var updateCourse model.Course
		updateCourse.ID = id
		json.NewDecoder(r.Body).Decode(&updateCourse)
		c.courseService.UpdateCourse(updateCourse)
		json.NewEncoder(w).Encode(updateCourse)

}
func (c *courseConnector)DeleteCourse(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("User-Count", strconv.Itoa(c.userService.GetUsersCount()))
		values := mux.Vars(r)
		id, _ := uuid.FromString(values["id"])
		var deleteCourse model.Course
		deleteCourse.ID = id
		json.NewDecoder(r.Body).Decode(&deleteCourse)
		c.courseService.DeleteCourse(deleteCourse)
		json.NewEncoder(w).Encode(deleteCourse)
	
}

