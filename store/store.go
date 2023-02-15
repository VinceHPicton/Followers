package store

import "fmt"

type Store struct {
	LargestUserID   int
	Users           map[int]User
	LargestFollowID int
	Follows         map[int]Follow
}

type User struct {
	name string
}

type Follow struct {
	follower_id int
	target_id   int
}

func (s *Store) CreateUser(name string) error {
	newUserID := len(s.Users) + 1

	_, ok := s.Users[newUserID]
	if ok {
		return fmt.Errorf("Attempt to create new User with name: %v failed: new id: %v already exists!", name, newUserID)
	}

	s.Users[newUserID] = User{
		name: name,
	}
	s.LargestUserID = newUserID

	return nil
}
