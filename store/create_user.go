package store

import "fmt"

func (s *Store) CreateUser(name string) (int, error) {
	newUserID := len(s.Users) + 1

	_, ok := s.Users[newUserID]
	if ok {
		return -1, fmt.Errorf("Attempt to create new User with name: %v failed: new id: %v already exists!", name, newUserID)
	}

	s.Users[newUserID] = User{
		name: name,
	}
	s.LargestUserID = newUserID

	return newUserID, nil
}
