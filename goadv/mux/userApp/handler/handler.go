package handler
import(
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"time"
	"reflect"
	service"pass/service"
	"fmt"
	"log"
	"net/http"
	"pass/model"
	"encoding/json"
)
var secretKey = []byte("nehaaaaaa")
type Response struct {
	Token  string `json:"token"`
	Status string `json:"status"`
}

type Handler struct{
	userService     *service.UserService
}
func Newhandler(userService *service.UserService) *Handler {
	return &Handler{
		userService:     userService,
	}
}

func (h *Handler)GetTokenHandler(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Please pass the data as URL form encoded", http.StatusBadRequest)
		return
	}
	log.Println(r.Form)
	var login model.Login
	json.NewDecoder(r.Body).Decode(&login)
	email:=login.Email
	password:=login.Password
	// email := r.Form.Get("email")
	// password := r.Form.Get("password")
	log.Println("email ", email)
	log.Println("password", password)

	if userPassHash,ok:= h.userService.GetPasswordFromEmail(email);ok{
		
		if service.CheckPasswordHash(password,userPassHash){
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
			h.userService.Logger.Error().Err(err).Msg("Invalid Credentials")
			return
		}
	} else {
		http.Error(w, "User is not found", http.StatusNotFound)
		h.userService.Logger.Error().Err(err).Msg("User is not found")
		return
	}
}

func (h *Handler)ValidAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println(r.URL.Path)
		if r.URL.Path == "/login" {
			next.ServeHTTP(w, r)
			return
		}
		tokenString, err := request.HeaderExtractor{"access_token"}.ExtractToken(r)
		if err!=nil{
			h.userService.Logger.Error().Err(err).Msg("check the access token")
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
			h.userService.Logger.Error().Err(err).Msg("check the access token")
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
			h.userService.Logger.Error().Err(err).Msg("check the access token")
			return
		}
	})
}