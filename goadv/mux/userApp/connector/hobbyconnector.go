package connector

import (
	"encoding/json"
	"net/http"
	"pass/handler"
	"pass/model"
	"pass/service"
	"strconv"

	// "time"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

type hobbyConnector struct {
	handler      *handler.Handler
	hobbyService *service.HobbyService
}

func NewHobbyConnector(handler *handler.Handler, hobbyService *service.HobbyService) *hobbyConnector {
	return &hobbyConnector{
		handler:      handler,
		hobbyService: hobbyService,
	}
}
func (h *hobbyConnector) RegisterHobbyRoutes(authRoute *mux.Router, noAuthRoute *mux.Router) {
	noAuthRoute.HandleFunc("/hobbies", h.getHobbies).Methods("GET")

	noAuthRoute.HandleFunc("/hobbies/{id}", h.getHobbiesOfUser).Methods("GET")
	noAuthRoute.HandleFunc("/hobbies", h.getHobbies).Queries("limit", "{limit:[0-9]+}", "pageNo", "{pageNo:[0-9]+}").Methods("GET")
	// noAuthRoute.HandleFunc("/hobbies/{id}", h.updateHobby).Methods("PUT")
	authRoute.Use(h.handler.ValidAuth)
	authRoute.HandleFunc("/hobbies/{id}", h.updateHobby).Methods("PUT")
	authRoute.HandleFunc("/hobbies/{id}", h.getHobby).Methods("GET")
}

func (h *hobbyConnector) getHobbiesOfUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := uuid.FromString(params["userId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("ID invalid!")
		return
	}
	var hobby []model.Hobby
	err2 := h.hobbyService.GetHobbyByUserId(&hobby, id)
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Could on find Hobbies")
		return
	}
	json.NewEncoder(w).Encode(hobby)
}

func (h *hobbyConnector) getHobbies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	limit, _ := strconv.Atoi(r.FormValue("limit"))
	pageNo, _ := strconv.Atoi(r.FormValue("pageNo"))
	var hobbies []model.Hobby
	if limit < 0 && pageNo >= 2 {
		json.NewEncoder(w).Encode(hobbies)
		return
	}
	offset := limit * (pageNo - 1)
	h.hobbyService.GetHobbies(&hobbies, limit, offset)
	json.NewEncoder(w).Encode(hobbies)
}

func (h *hobbyConnector) getHobby(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := uuid.FromString(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode("incorrect id")
		return
	}
	if !(h.hobbyService.CheckHobby(id)) {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode("hobby doesn't exists")
		return
	}
	var hobby model.Hobby
	var str1 []string
	h.hobbyService.GetHobbyFromId(&hobby, id, str1)
	json.NewEncoder(w).Encode(hobby)

}

func (h *hobbyConnector) updateHobby(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := uuid.FromString(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode("incorrect id")
		return
	}

	if !(h.hobbyService.CheckHobby(id)) {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode("hobby doesn't exists")
		return
	}
	var updatedHobby model.Hobby
	updatedHobby.ID = id
	json.NewDecoder(r.Body).Decode(&updatedHobby)
	h.hobbyService.UpdateHobby(updatedHobby)
	json.NewEncoder(w).Encode("updatedHobby")

}

// func (h *hobbyConnector) addHobby(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	//w.Header().Set("Hobby-Count", strconv.Itoa(hc.hobbyService.GetHobbyCount()))
// 	var updatedHobby model.Hobby
// 	updatedHobby.ID = uuid.NewV4()
// 	updatedHobby.CreateBy = "yogesh"
// 	updatedHobby.CreateAt = time.Now()
// 	json.NewDecoder(r.Body).Decode(&updatedHobby)

// 	err2 := h.hobbyService.AddHobby(updatedHobby)
// 	if err2 != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode("Error in adding hobby")
// 		return
// 	}
// 	json.NewEncoder(w).Encode("Updated hobby")
// }
