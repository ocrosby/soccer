package user_service

import "database/sql"

type Application struct {
	Address string
	DB      *sql.DB
}

func NewApplication(address string, db *sql.DB) *Application {
	return &Application{
		Address: address,
		DB:      db,
	}
}
