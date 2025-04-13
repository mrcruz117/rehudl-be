package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/mrcruz117/rehudl-be/internal/database"
)

func TestCreateUser(t *testing.T) {

	tests := []struct {
		name  string
		email string
	}{
		{name: "John Doe", email: "john.doe@example.com"},
		{name: "Jane Doe", email: "jane.doe@example.com"},
		{name: "Jim Doe", email: "jim.doe@example.com"},
		{name: "Jill Doe", email: "jill.doe@example.com"},
		{name: "Jack Doe", email: "jack.doe@example.com"},
		{name: "Michael Cruz", email: "derp@example.com"},
		{name: "Luis Cardona", email: "derp2@example.com"},
	}

	db, err := InitSQLiteDB("./rehudl.db")
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	queries := database.New(db)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			user, err := queries.CreateUser(context.Background(), database.CreateUserParams{
				Name:  test.name,
				Email: test.email,
			})
			if err != nil {
				t.Fatalf("Failed to create user: %v", err)
			}

			fmt.Println(user)
		})
	}

	users, err := queries.GetAllUsers(context.Background())
	if err != nil {
		t.Fatalf("Failed to get all users: %v", err)
	}

	fmt.Println(users)
}
