package main

import "fmt"

type (
	PostgresDb struct{}
	MockDb     struct{}
)

func InsertPlayerItemPostgres(db *PostgresDb) {
	db.Insert()
}

func (p *PostgresDb) Insert() {
	// Need real connection
	fmt.Println("Postgres Insert")
}

func InsertPlayerItemMock(db *MockDb) {
	db.Insert()
}

func (m *MockDb) Insert() {
	fmt.Println("Mock Insert")
}

func main() {
	postgresDb := &PostgresDb{}
	InsertPlayerItemPostgres(postgresDb)

	mockDb := &MockDb{}
	InsertPlayerItemMock(mockDb)
}
