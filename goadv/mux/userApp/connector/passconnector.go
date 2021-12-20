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
func RegisterPassportRoutes(db *gorm.DB,router *mux.Router) {
	router.HandleFunc("/passports", GetPassports(db)).Methods("GET")
	router.HandleFunc("/allpassports", GetAllPassports(db)).Methods("GET")
	router.HandleFunc("/passports", GetAllPassports(db)).Queries("limit", "{limit:[0-9]+}", "pageNo", "{pageNo:[0-9]+}").Methods("GET")
	router.HandleFunc("/passports/{id}", UpdatePassport(db)).Methods("PUT")
	router.HandleFunc("/passports/{id}", GetPassportFromId(db)).Methods("GET")
	router.HandleFunc("/users/{id}/passports",GetPassportByUserId(db)).Methods("GET")
}


// GetUserInfo(env *Env)http.HandlerFunc{
// 	return func(w http.ResponseWriter, r *http.Request){
func GetPassports(db *gorm.DB)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("User-Count", strconv.Itoa(userService.GetUsersCount(db)))
	var passports []model.Passport
	var str []string
	passportService.GetPassports(db,&passports, str)
	json.NewEncoder(w).Encode(passports)
	}
}
func GetAllPassports(db *gorm.DB)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("User-Count", strconv.Itoa(userService.GetUsersCount(db)))
	limit, _ := strconv.Atoi(r.FormValue("limit"))
	pageNo, _ := strconv.Atoi(r.FormValue("pageNo"))
	offset := limit * (pageNo - 1)
	fmt.Println(limit, pageNo)
	var passports []model.Passport
	passportService.GetAllPassports(db,&passports, limit, offset)
	json.NewEncoder(w).Encode(passports)
	}
}
func GetPassportFromId(db *gorm.DB)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("User-Count", strconv.Itoa(userService.GetUsersCount(db)))
		values := mux.Vars(r)
		id, _ := uuid.FromString(values["id"])
		var passport model.Passport
		passport.ID=id
		var str1 []string
		passportService.GetPassportFromId(db,&passport,id,str1)
		json.NewEncoder(w).Encode(passport)

	}
}

func UpdatePassport(db *gorm.DB)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("User-Count", strconv.Itoa(userService.GetUsersCount(db)))
		values := mux.Vars(r)
		id, _ := uuid.FromString(values["id"])
		var updatePassport model.Passport
		updatePassport.ID = id
		json.NewDecoder(r.Body).Decode(&updatePassport)
		passportService.UpdatePassport(db,updatePassport)
		json.NewEncoder(w).Encode(updatePassport)
	}
}

func GetPassportByUserId(db *gorm.DB)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("User-Count", strconv.Itoa(userService.GetUsersCount(db)))
	params := mux.Vars(r)
	id, _ := uuid.FromString(params["id"])
	var passport model.Passport
	passportService.GetPassportByUserId(db,&passport,id)
	json.NewEncoder(w).Encode(passport)
	}
}

func DeletePassport(db *gorm.DB)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("User-Count", strconv.Itoa(userService.GetUsersCount(db)))
		values := mux.Vars(r)
		id, _ := uuid.FromString(values["id"])
		// var deletePass model.Passport
		// deletePass.ID = id
		passportService.DeletePassport(db,id)
		json.NewEncoder(w).Encode("delete passport done")
	}
}