package main

import (
	"medical-records/config"
	"medical-records/models"
	"net/http"
)

func main() {
	db, err := models.NewDB("postgres://guest:guest@localhost/medicalrecords?sslmode=disable")
	if err != nil {
		config.Logger.Panic(err)
	}

	env := &Env{db}

	http.HandleFunc("/", env.MedicsIndex)
	http.ListenAndServe(":3000", nil)
}
