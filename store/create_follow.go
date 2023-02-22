package store

import "fmt"

func (s *Store) CreateFollow(followerID int, targetID int) error {
	newFollowID := len(s.Follows) + 1

	_, ok := s.Follows[newFollowID]
	if ok {
		return fmt.Errorf("Attempt to create new Follow failed: new id: %v already exists!", newFollowID)
	}

	s.Follows[newFollowID] = Follow{
		follower_id: followerID,
		target_id:   targetID,
	}
	s.LargestFollowID = newFollowID

	return nil
}
