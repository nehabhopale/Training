package handler

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"

	//"reflect"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"pass/model"
	service "pass/service"
)

var secretKey = []byte("nehaaaaaa")

type Response struct {
	Token  string `json:"token"`
	Status string `json:"status"`
}

type Handler struct {
	userService *service.UserService
}

func Newhandler(userService *service.UserService) *Handler {
	return &Handler{
		userService: userService,
	}
}
func (h *Handler) GetUserFromEmail(email string) (model.User, bool) {
	var user model.User
	_, err := h.userService.GetUserFromEmail(&user, email)
	if err != nil {
		return model.User{}, false
	}
	return user, true

}
func (h *Handler) GetTokenHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Please pass the data as URL form encoded", http.StatusBadRequest)
		return
	}
	log.Println(r.Form)
	var login model.Login
	json.NewDecoder(r.Body).Decode(&login)
	email := login.Email
	password := login.Password
	// email := r.Form.Get("email")
	// password := r.Form.Get("password")
	log.Println("email ", email)
	log.Println("password", password)

	user, ok := h.GetUserFromEmail(email)
	if ok {
		password = password + user.FirstName + user.LastName
		// fmt.Println("///////user////", password)
	}

	if userPassHash, ok := h.userService.GetPasswordFromEmail(email); ok {
		// fmt.Println("insideee okkkkkkkkkk")

		if service.CheckPasswordHash(password, userPassHash) {
			// fmt.Println("inside checkingggg")
			// Create a claims map
			claims := jwt.MapClaims{
				"email":    email,
				"IssuedAt": time.Now().Unix(),
			}
			claims["exp"] = time.Now().Add(time.Minute * 5).Unix()
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
			h.userService.Logger.Error().Err(err).Msg("Invalid Credentials")
			return
		}
	} else {
		http.Error(w, "User is not found", http.StatusBadRequest)
		h.userService.Logger.Error().Err(err).Msg("User is not found")
		return
	}
}

type AuthToken struct {
	Token string `json:"token"`
}

func (h *Handler) CheckToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tokenInfo AuthToken
	json.NewDecoder(r.Body).Decode(&tokenInfo)

	token, err := jwt.Parse(tokenInfo.Token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return secretKey, nil
	})

	if err != nil {
		//http.Error(w, "Forbidden", http.StatusForbidden)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid token")
		return
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("valid token")
	} else {
		//http.Error(w, "Forbidden", http.StatusForbidden)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("invalid tokenn")
		return
	}
}

func (h *Handler) ValidAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println(r.URL.Path)
		if r.URL.Path == "/login" {
			next.ServeHTTP(w, r)
			return
		}
		tokenString, err := request.HeaderExtractor{"access_token"}.ExtractToken(r)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Access Denied; Please check the access token"))
			// h.userService.Logger.Error().Err(err).Msg("check the access token")
			fmt.Fprintf(w, err.Error())
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
		//log.Println(reflect.TypeOf(token))
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, err.Error())
			return
			// h.userService.Logger.Error().Err(err).Msg("check the access token")
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// If token is valid
			// We found the token in our map
			//log.Printf("Authenticated user %s\n", claims)
			h.userService.Logger.Error().Interface("claims", claims).Msg("Authenticated user")
			// Pass down the request to the next middleware (or final handler)
			next.ServeHTTP(w, r)

		} else {
			w.WriteHeader(http.StatusForbidden)
			fmt.Println("*************")
			w.Write([]byte("check the access token"))
			h.userService.Logger.Error().Err(err).Msg("check the access token")
			return
		}
	})
}
