package main


import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"sync"
	
)
type Info1 struct{
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

var jokes map[int]Info1
var wg=sync.WaitGroup{}

func getJoke(i int ){
	
	resp, err := http.Get("https://api.chucknorris.io/jokes/random")
    var chuck Info1

    err = json.NewDecoder(resp.Body).Decode(&chuck)

    if err != nil {

        fmt.Println(err)

    }
	jokes[i]=chuck
    fmt.Println(jokes)
	wg.Done()
}
var mapp sync.Map
func getJoke1(i int){
	resp, err := http.Get("https://api.chucknorris.io/jokes/random")
    var chuck Info1

    err = json.NewDecoder(resp.Body).Decode(&chuck)

    if err != nil {

        fmt.Println(err)

    }
	//jokes[i]=chuck
    //fmt.Println(jokes)
	mapp.Store(i,chuck)
	wg.Done()

}
func main() {
	jokes=make(map[int]Info1,25)
    now := time.Now()
	
    defer func() {
		fmt.Println(now)
        fmt.Println(time.Since(now))

    }()
	wg.Add(25)
    // for i := 0; i < 25; i++ {

    //     go getJoke(i)
    // }
	var i int
	for i = 0; i < 25; i++ {
		go getJoke1(i)
	}

	wg.Wait()
	fmt.Println(mapp)
}