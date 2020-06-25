package services

import (
	"deprimera/src/application"
	"deprimera/src/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetEquipos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	equipos := models.NewEquipo(1, 1, "leonel", nil)

	j, err := json.Marshal(equipos)
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(j)
		log.Println(string(j))
	}
}

func SaveEquipos(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	equipos := &models.Equipos{}
	json.NewDecoder(r.Body).Decode(equipos)

	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())

	}
	db.Create(equipos)

	log.Println(equipos)
	w.Write([]byte("insertado"))
}

func InfoEquipo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	equipos := &models.Equipos{}
	json.NewDecoder(r.Body).Decode(equipos)

	j, err := json.Marshal(equipos)
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(j)
		log.Println(string(j))
	}
}