package connector

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"pass/model"
	"pass/service"
	"pass/handler"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)
type courseConnector struct{
	handler  *handler.Handler
	courseService *service.CourseService
}
func NewCourseConnector(handler *handler.Handler, courseService *service.CourseService) *courseConnector {
	return &courseConnector{
		handler:handler,
		courseService: courseService,
	}
}

func (c *courseConnector)RegisterCourseRoutes(authRoute *mux.Router,noAuthRoute *mux.Router) {
	noAuthRoute.HandleFunc("/courses", c.addCourse).Methods("POST")
	authRoute.Use(c.handler.ValidAuth)
	authRoute.HandleFunc("/courses", c.getAllCourses).Methods("GET")
	authRoute.HandleFunc("/courses",c.getAllCourses).Queries("limit", "{limit:[0-9]+}", "pageNo", "{pageNo:[0-9]+}").Methods("GET")
	authRoute.HandleFunc("/courses/{id}", c.updateCourse).Methods("PUT")
	authRoute.HandleFunc("/courses/{id}", c.getCourseFromId).Methods("GET")
	authRoute.HandleFunc("/courses/{id}", c.deleteCourse).Methods("DELETE")
}
func(c *courseConnector)addCourse(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var course model.Course
	json.NewDecoder(r.Body).Decode(&course)
	c.courseService.AddCourse(&course)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(course)
	
}

func (c *courseConnector)getAllCourses(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	limit, _ := strconv.Atoi(r.FormValue("limit"))
	pageNo, _ := strconv.Atoi(r.FormValue("pageNo"))
	offset := limit * (pageNo - 1)
	fmt.Println(limit, pageNo)
	var courses []model.Course
	c.courseService.GetAllCourses(&courses, limit, offset)
	json.NewEncoder(w).Encode(courses)
	
}
func (c *courseConnector)getCourseFromId(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		values := mux.Vars(r)
		id, err:= uuid.FromString(values["id"])
		if err!=nil{
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("incorrect id")
			return 
		}
		
		if !(c.courseService.CheckCourse(id)){
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("course doesn't exists")
			return 
		}
		var course model.Course
		course.ID=id
		var str1 []string
		c.courseService.GetCourseFromId(&course,id,str1)
		json.NewEncoder(w).Encode(course)

}

func (c *courseConnector)updateCourse(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		values := mux.Vars(r)
		id, err:= uuid.FromString(values["id"])
		if err!=nil{
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("incorrect id")
			return 
		}
		
		if !(c.courseService.CheckCourse(id)){
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("course doesn't exists")
			return 
		}
		var updateCourse model.Course
		updateCourse.ID = id
		json.NewDecoder(r.Body).Decode(&updateCourse)
		c.courseService.UpdateCourse(updateCourse)
		json.NewEncoder(w).Encode("updateCourse")

}
func (c *courseConnector)deleteCourse(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		values := mux.Vars(r)
		id, err := uuid.FromString(values["id"])
		if err!=nil{
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("incorrect id")
			return 
		}
		
		if !(c.courseService.CheckCourse(id)){
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("course doesn't exists")
			return 
		}
		var deleteCourse model.Course
		deleteCourse.ID = id
		json.NewDecoder(r.Body).Decode(&deleteCourse)
		c.courseService.DeleteCourse(deleteCourse)
		json.NewEncoder(w).Encode("deleteCourse")
	
}