package store

import "testing"

func TestCreateFollow(t *testing.T) {
	store := Start()

	followerName := "follower"
	targetName := "target"

	followerID, err := store.CreateUser(followerName)
	if err != nil {
		t.Errorf("Unexpected error: %v", err.Error())
	}

	targetID, err := store.CreateUser(targetName)
	if err != nil {
		t.Errorf("Unexpected error: %v", err.Error())
	}

	err = store.CreateFollow(followerID, targetID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err.Error())
	}

	createdFollow, ok := store.Follows[1]
	if !ok {
		t.Error("Unexpectedly failed to create the follow with id 1")
	}

	if createdFollow.follower_id != followerID {
		t.Errorf("Follower id was wrong, expected: %v, got: %v", followerID, createdFollow.follower_id)
	}
	if createdFollow.target_id != targetID {
		t.Errorf("targetID was wrong, expected: %v, got: %v", targetID, createdFollow.target_id)
	}
}
