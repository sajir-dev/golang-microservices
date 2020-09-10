package domain

// User is a struct to get exported, dont worry lint
type User struct {
	ID        uint64 `json: "message"`
	FirstName string `json: "first_name"`
	LastName  string `json: "last_name"`
	Email     string `json: "email"`
}
