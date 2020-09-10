package domain

import (
	"fmt"
	"net/http"

	"../utils"
)

var (
	users = map[int64]*User{
		123: {ID: 123, FirstName: "Aby", LastName: "Anil", Email: "aby@email.com"},
	}
)

// GetUser ...
func GetUser(userID int64) (*User, *utils.ApplicationError) {
	user := users[userID]
	if user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
