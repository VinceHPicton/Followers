package store

import "testing"

func TestCreateUser_Basic2UserCreation(t *testing.T) {
	store := Start()

	user1Name := "test"
	user1ExpectedID := 1
	user2Name := "test2"
	user2ExpectedID := 2

	user1ID, err := store.CreateUser(user1Name)
	if err != nil {
		t.Errorf("Unexpected error: %v", err.Error())
	}
	if user1ID != user1ExpectedID {
		t.Errorf("First UserID was not as expected: got: %v, expect: %v", user1ID, user1ExpectedID)
	}

	user2ID, err := store.CreateUser(user2Name)
	if err != nil {
		t.Errorf("Unexpected error: %v", err.Error())
	}
	if user2ID != user2ExpectedID {
		t.Errorf("Second UserID was not as expected: got: %v, expect: %v", user2ID, user2ExpectedID)
	}

	user1, ok := store.Users[user1ID]
	if !ok {
		t.Errorf("Failed to create first user")
	}
	if user1.name != user1Name {
		t.Errorf("User at id: %v was unexpected, user struct retrieved for the id: %+v", user1Name, user1)
	}

	user2, ok := store.Users[user2ID]
	if !ok {
		t.Errorf("Failed to create second user")
	}
	if user2.name != user2Name {
		t.Errorf("User at id: %v was unexpected, user struct retrieved for the id: %+v", user2Name, user2)
	}
}
