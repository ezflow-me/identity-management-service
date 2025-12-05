package domain

import "fmt"

func NewErrInvalidUserData(reason string) error {
	return fmt.Errorf("invalid user data: %s", reason)
}

func NewErrCannotSaveUser(reason string) error {
	return fmt.Errorf("cannot save user: %s", reason)
}

func NewErrCannotFindUserByID(id string, reason string) error {
	return fmt.Errorf("cannot find user by id %s: %s", id, reason)
}

func NewErrCannotDeleteUser(id string, reason string) error {
	return fmt.Errorf("cannot delete user with id %s: %s", id, reason)
}

func NewErrUserNotFound(id string) error {
	return fmt.Errorf("user with id %s not found", id)
}
