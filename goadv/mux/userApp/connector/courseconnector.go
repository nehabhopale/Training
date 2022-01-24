package connector

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pass/handler"
	"pass/model"
	"pass/service"
	"strconv"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

type courseConnector struct {
	handler       *handler.Handler
	courseService *service.CourseService
}

func NewCourseConnector(handler *handler.Handler, courseService *service.CourseService) *courseConnector {
	return &courseConnector{
		handler:       handler,
		courseService: courseService,
	}
}

func (c *courseConnector) RegisterCourseRoutes(authRoute *mux.Router, noAuthRoute *mux.Router) {
	noAuthRoute.HandleFunc("/courses", c.addCourse).Methods("POST")
	// noAuthRoute.HandleFunc("/courses", c.getAllCourses).Methods("GET")
	noAuthRoute.HandleFunc("/courses/{id}", c.deleteCourse).Methods("DELETE")
	noAuthRoute.HandleFunc("/courses", c.getAllCourses).Queries("limit", "{limit:[0-9]+}", "pageNo", "{pageNo:[0-9]+}").Methods("GET")
	// noAuthRoute.HandleFunc("/courses/{id}", c.updateCourse).Methods("PUT")
	authRoute.Use(c.handler.ValidAuth)
	authRoute.HandleFunc("/courses", c.getAllCourses).Methods("GET")
	// authRoute.HandleFunc("/courses",c.getAllCourses).Queries("limit", "{limit:[0-9]+}", "pageNo", "{pageNo:[0-9]+}").Methods("GET")
	authRoute.HandleFunc("/courses/{id}", c.updateCourse).Methods("PUT")
	authRoute.HandleFunc("/courses/{id}", c.getCourseFromId).Methods("GET")
	// authRoute.HandleFunc("/courses/{id}", c.deleteCourse).Methods("DELETE")
}
func (c *courseConnector) addCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var course model.Course
	json.NewDecoder(r.Body).Decode(&course)
	// courseName,_:=c.courseService.GetCourseFromName(&course,course.CourseName)
	// if courseName!=""{
	// 	w.WriteHeader(http.StatusForbidden)
	// 	json.NewEncoder(w).Encode("course already exist")
	// 	return
	// }
	err := c.courseService.AddCourse(&course)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode("error while adding course")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(course)

}

func (c *courseConnector) getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println("get all courses")
	limit, _ := strconv.Atoi(r.FormValue("limit"))
	pageNo, _ := strconv.Atoi(r.FormValue("pageNo"))
	offset := limit * (pageNo - 1)
	fmt.Println(limit, pageNo)
	var courses []model.Course
	c.courseService.GetAllCourses(&courses, limit, offset)
	json.NewEncoder(w).Encode(courses)

}
func (c *courseConnector) getCourseFromId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	values := mux.Vars(r)
	id, err := uuid.FromString(values["id"])
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode("incorrect id")
		return
	}

	if !(c.courseService.CheckCourse(id)) {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode("course doesn't exists")
		return
	}
	var course model.Course
	course.ID = id
	var str1 []string
	c.courseService.GetCourseFromId(&course, id, str1)
	json.NewEncoder(w).Encode(course)

}

func (c *courseConnector) updateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	values := mux.Vars(r)
	id, err := uuid.FromString(values["id"])
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode("incorrect id")
		return
	}

	if !(c.courseService.CheckCourse(id)) {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode("course doesn't exists")
		return
	}
	var updateCourse model.Course
	// var course model.Course
	// var c1 model.Course
	updateCourse.ID = id
	json.NewDecoder(r.Body).Decode(&updateCourse)
	// var str1 []string
	// c.courseService.GetCourseFromId(&course,id,str1)
	// course1,prize1,_:=c.courseService.GetCourseFromNamePrize(&c1,course.CourseName,updateCourse.Prize)
	// course2,prize2,_:=c.courseService.GetCourseFromNamePrize(&c1,updateCourse.CourseName,course.Prize)
	// course3,prize3,_:=c.courseService.GetCourseFromNamePrize(&c1,updateCourse.CourseName,updateCourse.Prize)
	// if( (course1!="" )||(course2!="" )||(course3!="" ))||((prize1!=-1)||(prize2!=-1)||(prize3!=-1)){
	// 	w.WriteHeader(http.StatusForbidden)
	// 	json.NewEncoder(w).Encode("course data already exist exists")
	// 	return
	// }
	err1 := c.courseService.UpdateCourse(updateCourse)
	if err1 != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode("error while updating course")
		return
	}
	json.NewEncoder(w).Encode("updateCourse")

}
func (c *courseConnector) deleteCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	values := mux.Vars(r)

	id, err := uuid.FromString(values["id"])

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode("incorrect id")
		return
	}

	if !(c.courseService.CheckCourse(id)) {
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
