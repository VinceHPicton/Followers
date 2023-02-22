package store

// GetUsersWhoFollow returns all the people that a given user follows
func (s *Store) GetUsersWhoFollow(userID int) []string {

	var userNamesFollowed []string

	for _, follow := range s.Follows {
		if follow.target_id == userID {
			userNamesFollowed = append(userNamesFollowed, s.Users[follow.follower_id].name)
		}
	}

	return userNamesFollowed
}
