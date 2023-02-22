package store

import "fmt"

func (s *Store) DeleteFollow(followerID int, targetID int) (int, error) {

	var followIDToDelete int

	for id, follow := range s.Follows {
		if follow.follower_id == followerID && follow.target_id == targetID {
			followIDToDelete = id
		}
	}

	if followIDToDelete == 0 {
		return -1, fmt.Errorf("Failed to find a valid follow for follower_id: %v, target_id: %v", followerID, targetID)
	}

	delete(s.Follows, followIDToDelete)

	return followIDToDelete, nil
}
