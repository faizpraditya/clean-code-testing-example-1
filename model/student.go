package model

import "time"

type Student struct {
	Id       int
	Name     string    `json:"name"`
	Gender   string    `json:"gender"`
	Age      int       `json:"age"`
	JoinDate time.Time `db:"join_date"`
	IdCard   string    `db:"id_card" json:"id_card"`
	Senior   bool      `json:"senior"`
}

type StudentJSON struct {
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Age      int    `json:"age"`
	JoinDate string `json:"join_date"`
	IdCard   string `json:"id_card"`
	Senior   bool   `json:"senior"`
}
