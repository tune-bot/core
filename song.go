package database

import (
	"os/exec"

	"github.com/google/uuid"
)

type Song struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}

func (s *Song) AddToPlaylist(playlistId string) error {
	songId := uuid.New().String()
	songExists := false
	_, err := db.Exec(`
		insert into song 
		(id, url) 
		values 
		(uuid_to_bin(?), ?);`,
		songId, s.Url)

	if err != nil {
		result, err := db.Query(`
			select bin_to_uuid(id) as id 
			from song 
			where url = ?;`,
			s.Url)

		if err == nil && result.Next() {
			songExists = true
			result.Scan(&songId)
		}
	}

	_, err = db.Exec(`
		insert into playlist_song 
		(id, playlist_id, song_id) 
		values 
		(uuid_to_bin(?), uuid_to_bin(?), uuid_to_bin(?));`,
		uuid.New().String(), playlistId, songId)

	s.Id = songId

	if !songExists {
		go download(*s)
	}

	return err
}

func (s Song) RemoveFromPlaylist(playlistId string) error {
	_, err := db.Exec(`
		delete from playlist_song 
		where playlist_id = uuid_to_bin(?)
		and song_id = uuid_to_bin(?);`,
		playlistId, s.Id)

	return err
}

func (s Song) FilePath() string {
	return "~/tune-bot/library/" + s.Id + ".mp3"
}

func download(s Song) {
	exec.Command("~/tune-bot/bin/download", "-o", "'"+s.FilePath()+"'", s.Url).Run()
}
