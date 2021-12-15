package main

import ("github.com/gorilla/mux"
"net/http"
"encoding/json"
"strconv"

)
type Student struct {
	ID int `json:"id"`
	Name string `json:"name"`
	
}
var students=[]Student{
	{ID :1,Name:"neha"},
	{ID :2,Name:"pooja"},

}
func getStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}
func createStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var student Student
	_ = json.NewDecoder(r.Body).Decode(&student)
	students = append(students, student)
	json.NewEncoder(w).Encode(student)
}
func deleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	for index, student := range students {
		if student.ID == id {
			students = append(students[:index], students[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(students)
}



func main(){
	router := mux.NewRouter()
	router.HandleFunc("/students", getStudents).Methods("GET")
	router.HandleFunc("/students", createStudent).Methods("POST")
	router.HandleFunc("/students/{id}", deleteStudent).Methods("DELETE")
	http.ListenAndServe(":8000", router)
}