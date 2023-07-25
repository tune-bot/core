package database

import "errors"

var (
	ErrDeviceNoUsers = errors.New("There are no users associated with this device")
	ErrInvalidLogin  = errors.New("Invalid login credentials")
	ErrNoDiscordUser = errors.New("There are no users associated with this Discord account")
	ErrUserNotFound  = errors.New("This user does not exist")
)
