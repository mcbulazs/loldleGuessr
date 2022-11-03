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

// Guesses from the frontend, 0 - red/wrong, 1 - amber/nearly, 2 - green/good
// In release years: 0 - lower, 1 - higher, 2 - exact
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
	switch r.Method {
	case http.MethodPost:
		sqlQuery = "SELECT `Name` FROM `Champs` WHERE 1=1 "
		var guesses []Guess

		if !DecodeRequest(w, r, &guesses) {
			return
		}

		for _, guess := range guesses {
			QueryBuilder(guess)
		}

		var possibleGuesses []string
		var nextGuess string
		rows, err := db.Query(sqlQuery)
		if err != nil {
			return
		}
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&nextGuess)
			possibleGuesses = append(possibleGuesses, nextGuess)
		}

		fmt.Println("\n" + sqlQuery)

		SendResponse(w, possibleGuesses)

	case http.MethodGet:
		var allChamps []string
		var nextChamp string

		rows, err := db.Query("SELECT `Name` FROM `Champs`")
		if err != nil {
			return
		}
		defer rows.Close()

		for rows.Next() {
			rows.Scan(&nextChamp)
			allChamps = append(allChamps, nextChamp)
		}

		SendResponse(w, allChamps)
	}
}

// Controller_Champs_Name returns the champ with the given name
func Controller_Champs_Name(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

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
