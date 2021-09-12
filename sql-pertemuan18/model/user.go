package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID      uint
	Name    string
	Email   string
	Address string
}

/*
*Contoh struct yg baik
	type User struct {
		ID           int       `db:"id" json:"id"`
		UserName     string    `db:"username" json:"username"`
		Email        string    `db:"email" json:"email"`
		CreatedAt    time.Time `db:"created_at" json:"created_at"`
		StatusID     uint8     `db:"status_id" json:"status_id"`
		Deleted      uint8     `db:"deleted" json:"deleted"`
		...
	}
*/
