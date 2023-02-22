package store

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

func Start() *Store {
	return &Store{
		Users:   make(map[int]User),
		Follows: make(map[int]Follow),
	}
}
