package main
import (
	"encoding/json"
	"fmt"
	"net/http"
	
)
type Info struct{
	// Name string `json:"name"`
	// Age  int    `json:"age"`
	Categories []	string `json:"categories"`
	Created_at  	string `json:"created_at"`
	Icon_url		string  `json:"icon_url"`
	Id 				string 	`json:"id"`
	Updated_at 		string `json:"updated_at"`
	Url        string   `json:"url"`
	Value      string   `json:"value"`
}


func main(){
	resp, err := http.Get("https://api.chucknorris.io/jokes/random")
    var jokes Info

    err = json.NewDecoder(resp.Body).Decode(&jokes)

    if err != nil {

        fmt.Println(err)

    }

    fmt.Println(jokes)
}