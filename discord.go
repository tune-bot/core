package database

import (
	"errors"

	"github.com/google/uuid"
)

var ErrNoDiscordUser = errors.New("There are no users associated with this Discord account")

type Discord struct {
	Id     string `json:"id"`
	UserId string `json:"userId"`
	Name   string `json:"name"`
}

func (d *Discord) Link() error {
	discordId := uuid.New().String()
	_, err := db.Exec(`
		insert into discord 
		(id, user_id, name) 
		values 
		(uuid_to_bin(?), uuid_to_bin(?), ?);`,
		discordId, d.UserId, d.Name)

	d.Id = discordId

	return err
}

func (d *Discord) GetUser() (User, error) {
	user := User{}

	if d.UserId == "" {
		result, err := db.Query("select bin_to_uuid(user_id) as user_id from discord where name = ?;", d.Name)

		if err != nil {
			return user, err
		}

		if !result.Next() {
			return user, ErrNoDiscordUser
		}
		result.Scan(&d.UserId)
	}

	user.Id = d.UserId
	err := user.getPlaylists()

	return user, err
}
