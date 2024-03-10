package main

import "fmt"

type (
	Database interface {
		Insert()
	}

	PostgresDb struct{}
	MockDb     struct{}
)

func NewPostgresDatabase() Database {
	return &PostgresDb{}
}

func (m *MockDb) Insert() {
	fmt.Println("Mock Insert")
}

func NewMockDatabase() Database {
	return &MockDb{}
}

func (p *PostgresDb) Insert() {
	// Need real connection
	fmt.Println("Postgres Insert")
}

func InsertPlayerItem(db Database) {
	db.Insert()
}

func main() {
	postgresDb := NewPostgresDatabase()
	InsertPlayerItem(postgresDb)

	mockDb := NewMockDatabase()
	InsertPlayerItem(mockDb)
}
