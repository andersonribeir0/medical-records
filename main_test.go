package main

import (
	"github.com/satori/go.uuid"
	"medical-records/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockDB struct{}

func (mdb *mockDB) AllMedics() ([]*models.Medic, error) {
	mds := make([]*models.Medic, 0)
	mds = append(mds, &models.Medic{Id: uuid.NewV1().String(), Name: "Rick", Crm: "123456789"})
	mds = append(mds, &models.Medic{Id: uuid.NewV1().String(), Name: "Jane", Crm: "987654321"})
	mds = append(mds, &models.Medic{Id: uuid.NewV1().String(), Name: "Jhon", Crm: "124578963"})
	return mds, nil
}

func TestMedicsIndex(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	env := &Env{new(mockDB)}
	router := setupRouter(env)
	router.ServeHTTP(rec,req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Invalid http code response: %v", rec.Code)
	}
}
