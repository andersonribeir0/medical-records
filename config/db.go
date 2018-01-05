package config

var Dbconn string

func init() {
	Dbconn = "postgres://guest:guest@localhost/medicalrecords?sslmode=disable"
}
