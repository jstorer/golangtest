package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
)

type catItem struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   string `json:"age"`
}

var catDB []catItem

func jsonResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func getAllCats(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w,http.StatusOK,catDB)
}

func createCat(w http.ResponseWriter, r *http.Request){
	var cat catItem
	_ = json.NewDecoder(r.Body).Decode(&cat)
	catDB = append(catDB,cat)
	jsonResponse(w, http.StatusCreated,cat)
}

func deleteCat(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	for index, item := range catDB{
		if item.ID == params["id"]{
			catDB = append(catDB[:index], catDB[index+1:]...)
			jsonResponse(w, http.StatusOK, catDB)
		}
	}
}

func main() {
	catDB = append(catDB, catItem{ID:"1", Name: "Felix", Breed: "Tabby", Age: "5"})
	catDB = append(catDB, catItem{ID:"2", Name: "Garfield", Breed: "Tiger", Age: "20"})
	router := mux.NewRouter()
	router.HandleFunc("/cats", getAllCats).Methods("GET")
	router.HandleFunc("/cats", createCat).Methods("POST")
	router.HandleFunc("/cats/{id}", deleteCat).Methods("DELETE")


	http.ListenAndServe(":8000", router)
}
