package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Champ struct {
	Name        string
	Gender      string
	Position    string
	Species     string
	Resource    string
	RangeType   string
	Region      string
	ReleaseYear int
}

type Guess struct {
	Name        string
	Gender      byte
	Position    byte
	Species     byte
	Resource    byte
	RangeType   byte
	Region      byte
	ReleaseYear byte
}

func Controller_Champs(w http.ResponseWriter, r *http.Request) {

}

// Controller_Champs_Name returns the champ with the given name
func Controller_Champs_Name(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	fmt.Println(name)

	result, err := GetChamp(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	SendResponse(w, result, "champ")
}

func GetChamp(name string) (*Champ, error) {
	champ := &Champ{}

	rows, err := db.Query("SELECT `Name`, `Gender`, `Position`, `Species`, `Resource`, `RangeType`, `Region`, `ReleaseYear` FROM `Champs` WHERE `Name` = ?", name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&champ.Name, &champ.Gender, &champ.Position, &champ.Species, &champ.Resource, &champ.RangeType, &champ.Region, &champ.ReleaseYear)
	} else {
		return nil, errors.New("champ not found")
	}

	return champ, nil
}
