package connector

import (
	"encoding/json"
	"os"
	"fmt"
	"net/http"
	"strconv"
	"log"
	"io/ioutil"
	"pass/model"
	repo"pass/repository"
	services"pass/services"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	//bcrypt"golang.org/x/crypto/bcrypt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"time"
	"reflect"
	"github.com/rs/zerolog"
)

var repo1 repo.Repository
var userService *services.UserService
var passportService *services.PassportService
var courseService *services.CourseService
var hobbyService *services.HobbyService
var db *gorm.DB
var secretKey = []byte("nehaaaaaa")
//var users = map[string]string{"naren": "passme", "admin": "password"}
// Response is a representation of JSON response for JWT
type Response struct {
	Token  string `json:"token"`
	Status string `json:"status"`
}
func Connect(dB *gorm.DB) {

	tempFile,_:= ioutil.TempFile(os.TempDir(), "deleteme")
	logger := zerolog.New(tempFile).With().Logger()
	db = dB
	repo1 = repo.NewRepository()
	userService = services.NewUserService(repo1,db,&logger)
	passportService = services.NewPassportService(repo1,&logger)
	courseService=services.NewCourseService(repo1,&logger)
	hobbyService=services.NewHobbyService(repo1,&logger)
	
}
func  RegisterUserRoutes(db *gorm.DB,authRout *mux.Router,nonAuthRoute *mux.Router) {
	nonAuthRoute.HandleFunc("/users/{id}", GetUserFromId(db)).Methods("GET")
	//fmt.Println("inside user route")
	//authRout.Use(ValidAuth)
	authRout.HandleFunc("/login", GetTokenHandler).Methods("GET")
	authRout.HandleFunc("/users", GetUsers(db)).Methods("GET")
	authRout.HandleFunc("/users",GetUsersWithPagination(db)).Methods("GET")
	authRout.HandleFunc("/users", GetUsersWithPagination(db)).Queries("limit", "{limit:[0-9]+}", "pageNo", "{pageNo:[0-9]+}").Methods("GET")
	authRout.HandleFunc("/users", AddUser(db)).Methods("POST")
	authRout.HandleFunc("/users/{id}", UpdateUser(db)).Methods("PUT")
	//authRout.HandleFunc("/users/{id}", GetUserFromId(db)).Methods("GET")
	authRout.HandleFunc("/users/{id}", DeleteUser(db)).Methods("DELETE")
}

// LoginHandler validates the user credentials
func GetTokenHandler(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Please pass the data as URL form encoded", http.StatusBadRequest)
		return
	}
	log.Println(r.Form)
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	log.Println("email ", email)
	log.Println("password", password)

	if userPassHash,ok:= userService.GetPasswordFromEmail(email);ok{
		
		if services.CheckPasswordHash(password,userPassHash){
			// Create a claims map
			claims := jwt.MapClaims{
				"email": email,
				"IssuedAt": time.Now().Unix(),
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString(secretKey)
			if err != nil {
				w.WriteHeader(http.StatusBadGateway)
				w.Write([]byte(err.Error()))
			}
			response := Response{Token: tokenString, Status: "success"}
			responseJSON, _ := json.Marshal(response)
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write(responseJSON)

		} else {
			http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
			userService.Logger.Error().Err(err).Msg("Invalid Credentials")
			return
		}
	} else {
		http.Error(w, "User is not found", http.StatusNotFound)
		userService.Logger.Error().Err(err).Msg("User is not found")
		return
	}
}

func ValidAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println(r.URL.Path)
		if r.URL.Path == "/login" {
			next.ServeHTTP(w, r)
			return
		}
		tokenString, err := request.HeaderExtractor{"access_token"}.ExtractToken(r)
		if err!=nil{
			userService.Logger.Error().Err(err).Msg("check the access token")
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return secretKey, nil
		})
		log.Println(reflect.TypeOf(token))
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			userService.Logger.Error().Err(err).Msg("check the access token")
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// If token is valid
			// We found the token in our map
			log.Printf("Authenticated user %s\n", claims)

			// Pass down the request to the next middleware (or final handler)
			next.ServeHTTP(w, r)

		} else {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("check the access token"))
			userService.Logger.Error().Err(err).Msg("check the access token")
			return
		}
	})
}


func GetUsersWithPagination(db *gorm.DB)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("User-Count", strconv.Itoa(userService.GetUsersCount(db)))
	limit, _ := strconv.Atoi(r.FormValue("limit"))
	pageNo, _ := strconv.Atoi(r.FormValue("pageNo"))
	offset := limit * (pageNo - 1)
	//fmt.Println(limit, pageNo)
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
	// pass,_:=services.HashPassword("b4567")
	var user model.User
	//user:=model.NewUser("neha","B","bneha@123",pass)
	json.NewDecoder(r.Body).Decode(&user)
	pass,_:=services.HashPassword(user.Password)
	user.Password=pass
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
		str1 :=[]string{"Passport","Courses","Hobbies"}
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
		var pass model.Passport
		if updateUser.Passport == pass {
			var passport model.Passport
			passportService.GetPassportByUserId(db,&passport, id)
			passportService.DeletePassport(db,passport.ID)
		}
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

