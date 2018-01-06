package handlers

import (
	. "medical-records/config"
	"medical-records/models"
	"net/http"
)

type Env struct {
	DB models.Datastore
}

func (env *Env) MedicsIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	mds, err := env.DB.AllMedics()
	if err != nil {
		Logger.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	if err := TPL.ExecuteTemplate(w, "medics.gohtml", mds); err != nil {
		Logger.Println(err)
	}
}
