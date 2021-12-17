package connector

import (
	"encoding/json"
	//"fmt"
	"net/http"
	"strconv"

	"pass/model"
	//repo"pass/repository"
//	services"pass/services"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)


// GetUserInfo(env *Env)http.HandlerFunc{
// 	return func(w http.ResponseWriter, r *http.Request){
func GetPassorts(db *gorm.DB)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("User-Count", strconv.Itoa(userService.GetUsersCount(db)))
	var passports []model.Passport
	 var str []string
	passportService.GetPassports(db,&passports, str)
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
		passport.PassId=id
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
		id, _ := uuid.FromString(values["pass_id"])
		var updatePassport model.Passport
		updatePassport.PassId = id
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