package connector

import (
	"encoding/json"
	"net/http"
	"strconv"
	"pass/model"
	services"pass/services"
	handler"pass/handler"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"

)

type userConnector struct{
	handler  *handler.Handler
	userService     *services.UserService
	passportService *services.PassportService
}
func NewUserConnector(handler *handler.Handler,userService *services.UserService, passportService *services.PassportService) *userConnector {
	return &userConnector{
		handler:handler,
		userService:     userService,
		passportService: passportService,
	}
}

var secretKey = []byte("nehaaaaaa")

func (u *userConnector) RegisterUserRoutes(authRoute *mux.Router,noAuthRoute *mux.Router) {
	noAuthRoute.HandleFunc("/users/{id}", u.GetUserFromId).Methods("GET")
	noAuthRoute.HandleFunc("/users", u.AddUser).Methods("POST")
	authRoute.Use(u.handler.ValidAuth)
	authRoute.HandleFunc("/users", u.GetUsers).Methods("GET")
	authRoute.HandleFunc("/users",u.GetUsersWithPagination).Methods("GET")
	authRoute.HandleFunc("/users", u.GetUsersWithPagination).Queries("limit", "{limit:[0-9]+}", "pageNo", "{pageNo:[0-9]+}").Methods("GET")
	authRoute.HandleFunc("/users/{id}", u.UpdateUser).Methods("PUT")
	authRoute.HandleFunc("/users/{id}", u.DeleteUser).Methods("DELETE")
}

func(u *userConnector)  GetUsersWithPagination(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("User-Count", strconv.Itoa(u.userService.GetUsersCount()))
	limit, _ := strconv.Atoi(r.FormValue("limit"))
	pageNo, _ := strconv.Atoi(r.FormValue("pageNo"))
	offset := limit * (pageNo - 1)
	var users []model.User
	u.userService.GetAllUsers(&users, limit, offset)
	json.NewEncoder(w).Encode(users)
}
func (u *userConnector)  GetUsers(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("User-Count", strconv.Itoa(u.userService.GetUsersCount()))
	var users []model.User
	str :=[]string{"Passport"}
	u.userService.GetUsers(&users, str)
	json.NewEncoder(w).Encode(users)
	
}
func(u *userConnector)  AddUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("User-Count", strconv.Itoa(u.userService.GetUsersCount()))
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	pass,_:=services.HashPassword(user.Password)
	user.Password=pass
	u.userService.AddUser(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func(u *userConnector)  GetUserFromId(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("User-Count", strconv.Itoa(u.userService.GetUsersCount()))
		values := mux.Vars(r)
		id, _ := uuid.FromString(values["id"])
		var user model.User
		str1 :=[]string{"Passport","Courses","Hobbies"}
		u.userService.GetUserFromId(&user,id,str1)
		json.NewEncoder(w).Encode(user)

	
}

func (u *userConnector)  UpdateUser(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("User-Count", strconv.Itoa(u.userService.GetUsersCount()))
		values := mux.Vars(r)
		id, _ := uuid.FromString(values["id"])
		var updateUser model.User
		updateUser.ID = id
		var pass model.Passport
		if updateUser.Passport == pass {
			var passport model.Passport
			u.passportService.GetPassportByUserId(&passport, id)
			u.passportService.DeletePassport(passport.ID)
		}
		json.NewDecoder(r.Body).Decode(&updateUser)
		u.userService.UpdateUser(updateUser)
		json.NewEncoder(w).Encode(updateUser)
	
}
func(u *userConnector)  DeleteUser(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("User-Count", strconv.Itoa(u.userService.GetUsersCount()))
		values := mux.Vars(r)
		id, _ := uuid.FromString(values["id"])
		var deleteUser model.User
		deleteUser.ID = id
		json.NewDecoder(r.Body).Decode(&deleteUser)
		u.userService.DeleteUser(deleteUser)
		json.NewEncoder(w).Encode(deleteUser)
	
}

