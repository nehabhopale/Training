package connector

import (
	"encoding/json"
	"net/http"
	"strconv"
	"pass/model"
	"pass/handler"
	"pass/service"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

type hobbyConnector struct{
	handler  *handler.Handler
	hobbyService *service.HobbyService
}
func NewHobbyConnector(handler *handler.Handler, hobbyService *service.HobbyService) *hobbyConnector {
	return &hobbyConnector{
		handler:handler,
		hobbyService: hobbyService,
	}
}
func(h *hobbyConnector) RegisterHobbyRoutes(authRoute *mux.Router,noAuthRoute *mux.Router) {
	noAuthRoute.HandleFunc("/hobbies", h.getHobbies).Methods("GET")
	noAuthRoute.HandleFunc("/hobbies", h.getHobbies).Queries("limit", "{limit:[0-9]+}", "pageNo", "{pageNo:[0-9]+}").Methods("GET")
	authRoute.Use(h.handler.ValidAuth)
	authRoute.HandleFunc("/hobbies/{id}", h.updateHobby).Methods("PUT")
	authRoute.HandleFunc("/hobbies/{id}",h.getHobby).Methods("GET")
}



func (h *hobbyConnector)getHobbies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	limit, _ := strconv.Atoi(r.FormValue("limit"))
	pageNo, _ := strconv.Atoi(r.FormValue("pageNo"))
	offset := limit * (pageNo - 1)
	var hobbies []model.Hobby
	h.hobbyService.GetHobbies(&hobbies, limit, offset)
	json.NewEncoder(w).Encode(hobbies)
	
}

func (h *hobbyConnector) getHobby(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := uuid.FromString(params["id"])
	var hobby model.Hobby
	var str1 []string
	h.hobbyService.GetHobbyFromId(&hobby, id,str1)
	json.NewEncoder(w).Encode(hobby)
	
}

func (h *hobbyConnector) updateHobby(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := uuid.FromString(params["id"])
	var updatedHobby model.Hobby
	updatedHobby.ID = id
	json.NewDecoder(r.Body).Decode(&updatedHobby)
	h.hobbyService.UpdateHobby(updatedHobby)
	json.NewEncoder(w).Encode(updatedHobby)
	
}