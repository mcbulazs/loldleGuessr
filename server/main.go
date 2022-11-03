package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func SendResponse(w http.ResponseWriter, i any, wrapper ...string) {
	data, err := json.Marshal(i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(wrapper) > 0 {
		data = append([]byte("{\""+wrapper[0]+"\":"), data...)
		data = append(data, []byte("}")...)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func DecodeRequest(w http.ResponseWriter, r *http.Request, i any) bool {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return false
	}
	return true
}

var db *sql.DB
var sqlQuery string

func main() {
	sqlQuery = "SELECT `Name` FROM `Champs` WHERE 1=1 "

	var err error
	db, err = sql.Open("mysql", "root:cica123@tcp(localhost:3306)/loldle")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MySQL")
	defer db.Close()

	mux := mux.NewRouter()
	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "X-Content-Type-Options"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "DELETE", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	mux.HandleFunc("/champs", Controller_Champs).Methods("GET", "POST")
	mux.HandleFunc("/champs/{name}", Controller_Champs_Name).Methods("GET")

	http.ListenAndServe(":7777", handlers.CORS(header, methods, origins)(mux))
}
