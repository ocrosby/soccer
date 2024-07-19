package database

import (
	"testing"
)

func TestNewSettings(t *testing.T) {
	host := "localhost"
	port := "5432"
	user := "user"
	password := "password"
	dbName := "testdb"
	sslMode := "disable"

	settings := NewSettings(host, port, user, password, dbName, sslMode)

	if settings.Host != host || settings.Port != port || settings.User != user || settings.Password != password || settings.DBName != dbName || settings.SSLMode != sslMode {
		t.Errorf("NewSettings() did not correctly initialize the settings")
	}
}

func TestConnectionString(t *testing.T) {
	expectedConnectionString := "host=localhost port=5432 user=user password=password dbname=testdb sslmode=disable"
	settings := &Settings{
		Host:     "localhost",
		Port:     "5432",
		User:     "user",
		Password: "password",
		DBName:   "testdb",
		SSLMode:  "disable",
	}

	if connStr := settings.ConnectionString(); connStr != expectedConnectionString {
		t.Errorf("ConnectionString() returned %v, want %v", connStr, expectedConnectionString)
	}
}
