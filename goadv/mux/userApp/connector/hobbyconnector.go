package connector

import (
	"encoding/json"
	"net/http"
	"strconv"
	"pass/model"
	"pass/handler"
	"pass/services"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

type hobbyConnector struct{
	handler  *handler.Handler
	userService     *services.UserService
	hobbyService *services.HobbyService
}
func NewHobbyConnector(handler *handler.Handler,userService *services.UserService, hobbyService *services.HobbyService) *hobbyConnector {
	return &hobbyConnector{
		handler:handler,
		userService: userService,
		hobbyService: hobbyService,
	}
}
func(h *hobbyConnector) RegisterHobbyRoutes(authRoute *mux.Router,noAuthRoute *mux.Router) {
	noAuthRoute.HandleFunc("/hobby", h.GetHobbies).Methods("GET")
	noAuthRoute.HandleFunc("/hobby", h.GetHobbies).Queries("limit", "{limit:[0-9]+}", "pageNo", "{pageNo:[0-9]+}").Methods("GET")
	authRoute.Use(h.handler.ValidAuth)
	authRoute.HandleFunc("/hobby/{id}", h.UpdateHobby).Methods("PUT")
	authRoute.HandleFunc("/hobby/{id}",h. GetHobby).Methods("GET")
}



func (h *hobbyConnector)GetHobbies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("User-Count", strconv.Itoa(h.userService.GetUsersCount()))
	limit, _ := strconv.Atoi(r.FormValue("limit"))
	pageNo, _ := strconv.Atoi(r.FormValue("pageNo"))
	offset := limit * (pageNo - 1)
	var hobbies []model.Hobby
	h.hobbyService.GetHobbies(&hobbies, limit, offset)
	json.NewEncoder(w).Encode(hobbies)
	
}

func (h *hobbyConnector) GetHobby(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("User-Count", strconv.Itoa(h.userService.GetUsersCount()))
	params := mux.Vars(r)
	id, _ := uuid.FromString(params["id"])
	var hobby model.Hobby
	var str1 []string
	h.hobbyService.GetHobbyFromId(&hobby, id,str1)
	json.NewEncoder(w).Encode(hobby)
	
}

func (h *hobbyConnector) UpdateHobby(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("User-Count", strconv.Itoa(h.userService.GetUsersCount()))
	params := mux.Vars(r)
	id, _ := uuid.FromString(params["id"])
	var updatedHobby model.Hobby
	updatedHobby.ID = id
	json.NewDecoder(r.Body).Decode(&updatedHobby)
	h.hobbyService.UpdateHobby(updatedHobby)
	json.NewEncoder(w).Encode(updatedHobby)
	
}