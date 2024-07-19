package database

import "fmt"

type Settingser interface {
	ConnectionString() string
}

// Settings is a struct that holds the database settings
type Settings struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// host=postgres port=%s user=user password=password dbname=testdb sslmode=disable

func NewSettings(host, port, user, password, dbName, sslMode string) *Settings {
	return &Settings{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DBName:   dbName,
		SSLMode:  sslMode,
	}
}

// ConnectionString returns the connection string for the database
func (s *Settings) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", s.Host, s.Port, s.User, s.Password, s.DBName, s.SSLMode)
}
