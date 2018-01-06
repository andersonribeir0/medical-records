package main

import (
	cfg "medical-records/config"
	"medical-records/models"
	"github.com/gin-gonic/gin"
	"fmt"
)
type Env struct {
	DB models.Datastore
}

func setupRouter(env *Env) *gin.Engine{
	r := gin.Default()
	r.GET("/",env.MedicsIndex)
	return r
}
func main() {
	db, err := models.NewDB("postgres://guest:guest@localhost/medicalrecords?sslmode=disable")
	if err != nil {
		cfg.Logger.Panic(err)
	}
	env := &Env{db}
	r := setupRouter(env)
	r.Run(":3000")
}

func (env *Env) MedicsIndex(c *gin.Context) {

	mds, err := env.DB.AllMedics()
	if err != nil {
		cfg.Logger.Println(err)
		c.String(500,fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	if err := cfg.TPL.ExecuteTemplate(c.Writer, "medics.gohtml", mds); err != nil {
		cfg.Logger.Println(err)
		c.String(500,fmt.Sprintf("get form err: %s", err.Error()))
	}
}