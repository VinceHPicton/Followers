package store

// GetFollowers returns all the people who a given user follows
func (s *Store) GetFollowers(userID int) []string {

	var followerUsernames []string

	for _, follow := range s.Follows {
		if follow.follower_id == userID {
			followerUsernames = append(followerUsernames, s.Users[follow.target_id].name)
		}
	}

	return followerUsernames
}
