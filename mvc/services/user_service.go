package services

import (
	"../domain"
	"../utils"
)

// GetUser ...
func GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userID)
}
