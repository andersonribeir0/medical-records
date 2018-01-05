package models

import (
	_ "github.com/lib/pq"
	"log"
)

type Medic struct {
	Id   string
	Name string
	Crm  string
}

func (db *DB) AllMedics() ([]*Medic, error) {
	log.Println("Starting to pull all medics")

	rows, err := db.Query("SELECT * FROM medics")
	if err != nil {
		return nil, err
	}
	log.Printf("Finished to pull all medics.")
	defer rows.Close()

	mds := make([]*Medic, 0)
	log.Println("Starting to scan all rows.")
	for rows.Next() {
		m := new(Medic)
		err := rows.Scan(&m.Id, &m.Name, &m.Crm)
		if err != nil {
			return nil, err
		}
		mds = append(mds, m)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	log.Printf("Finished to scan all rows. Total of %d rows.\n", len(mds))

	return mds, nil
}
