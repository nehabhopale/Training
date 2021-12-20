package connector

import (
	"encoding/json"
	//"fmt"
	"net/http"
	"strconv"
	"pass/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)
func RegisterHobbyRoutes(db *gorm.DB,router *mux.Router) {
	router.HandleFunc("/hobby", GetHobbies(db)).Methods("GET")
	router.HandleFunc("/hobby", GetHobbies(db)).Queries("limit", "{limit:[0-9]+}", "pageNo", "{pageNo:[0-9]+}").Methods("GET")
	router.HandleFunc("/hobby/{id}", UpdateHobby(db)).Methods("PUT")
	router.HandleFunc("/hobby/{id}", GetHobby(db)).Methods("GET")
}



func GetHobbies(db *gorm.DB)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("User-Count", strconv.Itoa(userService.GetUsersCount(db)))
	limit, _ := strconv.Atoi(r.FormValue("limit"))
	pageNo, _ := strconv.Atoi(r.FormValue("pageNo"))
	offset := limit * (pageNo - 1)
	var hobbies []model.Hobby
	hobbyService.GetHobbies(db,&hobbies, limit, offset)
	json.NewEncoder(w).Encode(hobbies)
	}
}

func  GetHobby(db *gorm.DB)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("User-Count", strconv.Itoa(userService.GetUsersCount(db)))
	params := mux.Vars(r)
	id, _ := uuid.FromString(params["id"])
	var hobby model.Hobby
	var str1 []string
	hobbyService.GetHobbyFromId(db,&hobby, id,str1)
	json.NewEncoder(w).Encode(hobby)
	}
}

func  UpdateHobby(db *gorm.DB)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("User-Count", strconv.Itoa(userService.GetUsersCount(db)))
	params := mux.Vars(r)
	id, _ := uuid.FromString(params["id"])
	var updatedHobby model.Hobby
	updatedHobby.ID = id
	json.NewDecoder(r.Body).Decode(&updatedHobby)
	hobbyService.UpdateHobby(db,updatedHobby)
	json.NewEncoder(w).Encode(updatedHobby)
	}
}