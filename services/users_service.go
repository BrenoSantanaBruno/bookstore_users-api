package services

import (
	"github.com/LinuxLoverCoder/bookstore_users-api/domain/users"
	"github.com/LinuxLoverCoder/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	return nil, nil
}
