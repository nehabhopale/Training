package connector

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"pass/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)
func RegisterCourseRoutes(db *gorm.DB,router *mux.Router) {
	router.HandleFunc("/course", GetAllCourses(db)).Methods("GET")
	router.HandleFunc("/course", GetAllCourses(db)).Queries("limit", "{limit:[0-9]+}", "pageNo", "{pageNo:[0-9]+}").Methods("GET")
	router.HandleFunc("/course", AddCourse(db)).Methods("POST")
	router.HandleFunc("/course/{id}", UpdateCourse(db)).Methods("PUT")
	router.HandleFunc("/course/{id}", GetCourseFromId(db)).Methods("GET")
	router.HandleFunc("/course/{id}", DeleteCourse(db)).Methods("DELETE")
}
func AddCourse(db *gorm.DB)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("User-Count", strconv.Itoa(userService.GetUsersCount(db)))
	//course:=model.NewCourse("python")
	// err2:=db.Debug().Model(course).Association("Users").Error
	// if err2!=nil{
	// 	fmt.Println("error in association------>",err2)
	// }
	var course model.Course
	json.NewDecoder(r.Body).Decode(&course)
	courseService.AddCourse(db,&course)
	err2:=db.Debug().Model(course).Association("Users").Error
	if err2!=nil{
		fmt.Println("error in association------>",err2)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(course)
	}
}

func GetAllCourses(db *gorm.DB)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("User-Count", strconv.Itoa(userService.GetUsersCount(db)))
	limit, _ := strconv.Atoi(r.FormValue("limit"))
	pageNo, _ := strconv.Atoi(r.FormValue("pageNo"))
	offset := limit * (pageNo - 1)
	fmt.Println(limit, pageNo)
	var courses []model.Course
	courseService.GetAllCourses(db,&courses, limit, offset)
	json.NewEncoder(w).Encode(courses)
	}
}
func GetCourseFromId(db *gorm.DB)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("User-Count", strconv.Itoa(userService.GetUsersCount(db)))
		values := mux.Vars(r)
		id, _ := uuid.FromString(values["id"])
		var course model.Course
		course.ID=id
		var str1 []string
		courseService.GetCourseFromId(db,&course,id,str1)
		json.NewEncoder(w).Encode(course)

	}
}

func UpdateCourse(db *gorm.DB)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("User-Count", strconv.Itoa(userService.GetUsersCount(db)))
		values := mux.Vars(r)
		id, _ := uuid.FromString(values["id"])
		var updateCourse model.Course
		updateCourse.ID = id
		json.NewDecoder(r.Body).Decode(&updateCourse)
		courseService.UpdateCourse(db,updateCourse)
		json.NewEncoder(w).Encode(updateCourse)
	}
}
func DeleteCourse(db *gorm.DB)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("User-Count", strconv.Itoa(userService.GetUsersCount(db)))
		values := mux.Vars(r)
		id, _ := uuid.FromString(values["id"])
		var deleteCourse model.Course
		deleteCourse.ID = id
		json.NewDecoder(r.Body).Decode(&deleteCourse)
		courseService.DeleteCourse(db,deleteCourse)
		json.NewEncoder(w).Encode(deleteCourse)
	}
}

