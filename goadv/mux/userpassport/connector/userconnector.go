package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"pass/model"
	repo"pass/repository"
	services"pass/services"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

var repo1 repo.Repository
var userService *services.UserService
var passportService *services.PassportService
var db *gorm.DB

func Connect(dB *gorm.DB) {
	db = dB
	repo1 = repo.NewRepository()
	
	userService = services.NewUserService(repo1)

	passportService = services.NewPassportService(repo1)
	
}


func GetUsersWithPagination(db *gorm.DB)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("User-Count", strconv.Itoa(userService.GetUsersCount(db)))
	limit, _ := strconv.Atoi(r.FormValue("limit"))
	pageNo, _ := strconv.Atoi(r.FormValue("pageNo"))
	offset := limit * (pageNo - 1)
	fmt.Println(limit, pageNo)
	var users []model.User
	userService.GetAllUsers(db,&users, limit, offset)
	json.NewEncoder(w).Encode(users)
}}
func GetUsers(db *gorm.DB)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("User-Count", strconv.Itoa(userService.GetUsersCount(db)))
	var users []model.User
	str :=[]string{"Passport"}
	userService.GetUsers(db,&users, str)
	json.NewEncoder(w).Encode(users)
	}
}



func AddUser(db *gorm.DB)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("User-Count", strconv.Itoa(userService.GetUsersCount(db)))
	//var user model.User
	user:=model.NewUser("pooja")
	json.NewDecoder(r.Body).Decode(&user)
	userService.AddUser(db,&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
	}
}

func GetUserFromId(db *gorm.DB)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("User-Count", strconv.Itoa(userService.GetUsersCount(db)))
		values := mux.Vars(r)
		id, _ := uuid.FromString(values["id"])
		var user model.User
		str1 :=[]string{"Passport"}
		userService.GetUserFromId(db,&user,id,str1)
		json.NewEncoder(w).Encode(user)

	}
}

func UpdateUser(db *gorm.DB)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("User-Count", strconv.Itoa(userService.GetUsersCount(db)))
		values := mux.Vars(r)
		id, _ := uuid.FromString(values["id"])
		var updateUser model.User
		updateUser.ID = id
		json.NewDecoder(r.Body).Decode(&updateUser)
		userService.UpdateUser(db,updateUser)
		json.NewEncoder(w).Encode(updateUser)
	}
}
func DeleteUser(db *gorm.DB)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("User-Count", strconv.Itoa(userService.GetUsersCount(db)))
		values := mux.Vars(r)
		id, _ := uuid.FromString(values["id"])
		var deleteUser model.User
		deleteUser.ID = id
		json.NewDecoder(r.Body).Decode(&deleteUser)
		userService.DeleteUser(db,deleteUser)
		json.NewEncoder(w).Encode(deleteUser)
	}
}

// func GetPassportByUserId(db *gorm.DB)http.HandlerFunc{
// 	return func(w http.ResponseWriter, r *http.Request){
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("User-Count", strconv.Itoa(userService.GetUsersCount(db)))
// 	params := mux.Vars(r)
// 	id, _ := uuid.FromString(params["id"])
// 	var passport model.Passport
// 	passportService.GetPassportByUserId(db,&passport,id)
// 	json.NewEncoder(w).Encode(passport)
// }
// }