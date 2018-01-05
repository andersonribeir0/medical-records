package main

import (
	. "medical-records/config"
	"medical-records/handlers"
	"medical-records/models"
	"net/http"
)

func main() {
	db, err := models.NewDB("postgres://guest:guest@localhost/medicalrecords?sslmode=disable")
	if err != nil {
		Logger.Panic(err)
	}

	env := &handlers.Env{db}

	http.HandleFunc("/", env.MedicsIndex)
	http.ListenAndServe(":3000", nil)
}
