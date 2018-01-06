package main

import (
	cfg "medical-records/config"
	"medical-records/models"
	"net/http"
)

//Env Enviroment struct containing the repository interface
type Env struct {
	DB models.Datastore
}

//MedicsIndex Return all medics
func (env *Env) MedicsIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	mds, err := env.DB.AllMedics()
	if err != nil {
		cfg.Logger.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	if err := TPL.ExecuteTemplate(w, "medics.gohtml", mds); err != nil {
		cfg.Logger.Println(err)
	}
}
