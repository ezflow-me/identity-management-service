package domain

import (
	"time"
)

type User struct {
	id              string    `validate:"required"`
	name            string    `validate:"required,min=2,max=100"`
	email           string    `validate:"required,email"`
	displayLanguage string    `validate:"oneof=en es"`
	timeZone        string    `validate:"timezone"`
	createdAt       time.Time `validate:"required"`
}

func NewUser(id, name, email string, createdAt float64) (*User, error) {
	user := &User{
		id:              id,
		name:            name,
		email:           email,
		createdAt:       time.UnixMilli(int64(createdAt)),
		displayLanguage: "",
		timeZone:        "",
	}

	return user, nil
}

func (u *User) UpdateDisplayLanguage(displayLanguage string) error {
	u.displayLanguage = displayLanguage

	return nil
}

func (u *User) UpdateTimeZone(timeZone string) error {
	u.timeZone = timeZone

	return nil
}

func (u *User) ID() string {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Email() string {
	return u.email
}

func (u *User) DisplayLanguage() string {
	return u.displayLanguage
}

func (u *User) TimeZone() string {
	return u.timeZone
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func (u *User) ToPrimitives() map[string]interface{} {
	return map[string]interface{}{
		"id":               u.id,
		"name":             u.name,
		"email":            u.email,
		"time_zone":        u.timeZone,
		"display_language": u.displayLanguage,
		"created_at":       u.createdAt,
	}
}

func (u *User) PtrsForScan() []interface{} {
	return []interface{}{
		&u.id,
		&u.name,
		&u.email,
		&u.timeZone,
		&u.displayLanguage,
		&u.createdAt,
	}
}
