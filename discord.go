package core

import (
	"github.com/google/uuid"
)

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
	result, err := db.Query(`
		select bin_to_uuid(u.id), u.username
		from user as u
		inner join discord AS d on u.id = d.user_id
		where d.name = ?;`,
		d.Name)

	if err != nil {
		PrintError(err.Error())
		return user, ErrNoUser
	}

	if result.Next() {
		result.Scan(&user.Id, &user.Username)
		err = user.getPlaylists()
		return user, err
	}

	return user, ErrNoUser
}
