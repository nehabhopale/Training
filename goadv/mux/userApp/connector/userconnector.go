package connector

import (
	"encoding/json"
	"fmt"
	"net/http"
	handler "pass/handler"
	"pass/model"
	service "pass/service"
	"strconv"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

type UserConnector struct {
	handler         *handler.Handler
	userService     *service.UserService
	passportService *service.PassportService
	hobbyService    *service.HobbyService
}

func NewUserConnector(handler *handler.Handler, userService *service.UserService, passportService *service.PassportService, hobbyService *service.HobbyService) *UserConnector {
	return &UserConnector{
		handler:         handler,
		userService:     userService,
		passportService: passportService,
		hobbyService:    hobbyService,
	}
}

var secretKey = []byte("nehaaaaaa")

func (u *UserConnector) RegisterUserRoutes(authRoute *mux.Router, noAuthRoute *mux.Router) {
	noAuthRoute.HandleFunc("/users/{id}", u.GetUserFromId).Methods("GET")
	noAuthRoute.HandleFunc("/users", u.addUser).Methods("POST")
	authRoute.Use(u.handler.ValidAuth)
	authRoute.HandleFunc("/users", u.getUsers).Methods("GET")
	authRoute.HandleFunc("/users", u.getUsersWithPagination).Methods("GET")
	authRoute.HandleFunc("/users", u.getUsersWithPagination).Queries("limit", "{limit:[0-9]+}", "pageNo", "{pageNo:[0-9]+}").Methods("GET")
	authRoute.HandleFunc("/users/{id}", u.updateUser).Methods("PUT")
	authRoute.HandleFunc("/users/{id}", u.deleteUser).Methods("DELETE")
}

func (u *UserConnector) getUsersWithPagination(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("User-Count", strconv.Itoa(u.userService.GetUsersCount()))
	limit, _ := strconv.Atoi(r.FormValue("limit"))
	pageNo, _ := strconv.Atoi(r.FormValue("pageNo"))
	offset := limit * (pageNo - 1)
	var users []model.User
	u.userService.GetAllUsers(&users, limit, offset)
	json.NewEncoder(w).Encode(users)
}
func (u *UserConnector) getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("User-Count", strconv.Itoa(u.userService.GetUsersCount()))
	var users []model.User
	str := []string{"Passport"}
	u.userService.GetUsers(&users, str)
	json.NewEncoder(w).Encode(users)

}

func (u *UserConnector) addUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("User-Count", strconv.Itoa(u.userService.GetUsersCount()))
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	pass, _ := service.HashPassword(user.Password + user.FirstName + user.LastName)
	fmt.Println("pass while adding user", pass)
	user.Password = pass
	email, _ := u.userService.GetUserFromEmail(&user, user.Email)
	if email != "" {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode("email already exist")
		return
	}
	err := u.userService.AddUser(&user)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode("eror while adding user")
		return

	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (u *UserConnector) GetUserFromId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("User-Count", strconv.Itoa(u.userService.GetUsersCount()))
	values := mux.Vars(r)
	id, err := uuid.FromString(values["id"])
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode("incorrect id")
		return
	}

	if !(u.userService.CheckUser(id)) {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode("user doesn't exists")
		return
	}
	var user model.User
	str1 := []string{"Passport", "Courses", "Hobbies"}
	u.userService.GetUserFromId(&user, id, str1)
	json.NewEncoder(w).Encode(user)

}

func (u *UserConnector) updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("User-Count", strconv.Itoa(u.userService.GetUsersCount()))
	values := mux.Vars(r)
	id, err := uuid.FromString(values["id"])
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode("incorrect id")
		return
	}
	var updatedUser model.User
	json.NewDecoder(r.Body).Decode(&updatedUser)
	updatedUser.ID = id

	var pass model.Passport
	if updatedUser.Passport == pass {
		var passport model.Passport
		u.passportService.GetPassportByUserId(&passport, id)
		u.passportService.DeletePassport(passport.ID)
	}
	if !(u.userService.CheckUser(id)) {
		json.NewEncoder(w).Encode("user doesn't exists")
		return
	}

	u.userService.UpdateUser(&updatedUser)
	json.NewEncoder(w).Encode("updatedUser")

}
func (u *UserConnector) deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("User-Count", strconv.Itoa(u.userService.GetUsersCount()))
	values := mux.Vars(r)
	id, err := uuid.FromString(values["id"])
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode("incorrect id")
		return
	}

	if !(u.userService.CheckUser(id)) {
		json.NewEncoder(w).Encode("user doesn't exists")
		return
	}

	var user model.User
	str1 := []string{"Passport", "Courses", "Hobbies"}
	u.userService.GetUserFromId(&user, id, str1)
	fmt.Println(user)
	u.passportService.DeletePassport(user.Passport.ID)
	fmt.Println(user.Hobbies)
	for _, hobby := range user.Hobbies {
		u.hobbyService.DeleteHobby(hobby)
	}
	var deleteUser model.User
	deleteUser.ID = id
	json.NewDecoder(r.Body).Decode(&deleteUser)
	u.userService.DeleteUser(deleteUser)

	json.NewEncoder(w).Encode("deleteUser")

}
