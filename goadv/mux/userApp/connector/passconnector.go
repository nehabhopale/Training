package connector

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"pass/model"
	"pass/handler"
	"pass/service"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

type passConnector struct{
	handler  *handler.Handler
	userService     *service.UserService
	passportService *service.PassportService
}
func NewPassportConnector(handler *handler.Handler,userService *service.UserService, passService *service.PassportService) *passConnector {
	return &passConnector{
		handler:handler,
		userService: userService,
		passportService: passService,
	}
}
func(p *passConnector) RegisterPassportRoutes(authRoute *mux.Router,nonAuthRoute *mux.Router) {
	authRoute.Use(p.handler.ValidAuth)
	authRoute.HandleFunc("/passports", p.getPassports).Methods("GET")
	authRoute.HandleFunc("/allpassports", p.getAllPassports).Methods("GET")
	authRoute.HandleFunc("/passports", p.getAllPassports).Queries("limit", "{limit:[0-9]+}", "pageNo", "{pageNo:[0-9]+}").Methods("GET")
	authRoute.HandleFunc("/passports/{id}",p.updatePassport).Methods("PUT")
	authRoute.HandleFunc("/passports/{id}", p.getPassportFromId).Methods("GET")
	authRoute.HandleFunc("/users/{id}/passports",p.GetPassportByUserId).Methods("GET")
}
func (p *passConnector)getPassports(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("User-Count", strconv.Itoa(p.userService.GetUsersCount()))
	var passports []model.Passport
	var str []string
	p.passportService.GetPassports(&passports, str)
	json.NewEncoder(w).Encode(passports)
	
}
func(p *passConnector) getAllPassports(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("User-Count", strconv.Itoa(p.userService.GetUsersCount()))
	limit, _ := strconv.Atoi(r.FormValue("limit"))
	pageNo, _ := strconv.Atoi(r.FormValue("pageNo"))
	offset := limit * (pageNo - 1)
	fmt.Println(limit, pageNo)
	var passports []model.Passport
	p.passportService.GetAllPassports(&passports, limit, offset)
	json.NewEncoder(w).Encode(passports)
	
}
func(p *passConnector) getPassportFromId(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("User-Count", strconv.Itoa(p.userService.GetUsersCount()))
		values := mux.Vars(r)
		id, _ := uuid.FromString(values["id"])
		var passport model.Passport
		passport.ID=id
		var str1 []string
		p.passportService.GetPassportFromId(&passport,id,str1)
		json.NewEncoder(w).Encode(passport)
}

func(p *passConnector) updatePassport(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("User-Count", strconv.Itoa(p.userService.GetUsersCount()))
		values := mux.Vars(r)
		id, _ := uuid.FromString(values["id"])
		var updatePassport model.Passport
		updatePassport.ID = id
		json.NewDecoder(r.Body).Decode(&updatePassport)
		p.passportService.UpdatePassport(updatePassport)
		json.NewEncoder(w).Encode(updatePassport)
	
}

func (p *passConnector)GetPassportByUserId(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("User-Count", strconv.Itoa(p.userService.GetUsersCount()))
	params := mux.Vars(r)
	id, _ := uuid.FromString(params["id"])
	var passport model.Passport
	p.passportService.GetPassportByUserId(&passport,id)
	json.NewEncoder(w).Encode(passport)
	
}

func (p *passConnector)DeletePassport(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("User-Count", strconv.Itoa(p.userService.GetUsersCount()))
		values := mux.Vars(r)
		id, _ := uuid.FromString(values["id"])
		p.passportService.DeletePassport(id)
		json.NewEncoder(w).Encode("delete passport done")
	
}