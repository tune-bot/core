package core

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/google/uuid"
)

type Song struct {
	Id     string `json:"id"`
	Url    string `json:"url"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Album  string `json:"album"`
	Year   uint16 `json:"year"`
}

func (s *Song) AddToPlaylist(playlistId string) error {
	songId := uuid.New().String()
	songExists := false
	_, err := db.Exec(`
		insert into song 
		(id, url, title, artist, album, year) 
		values 
		(uuid_to_bin(?), ?, ?, ?, ?, ?);`,
		songId, s.Url, s.Title, s.Artist, s.Album, s.Year)

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
	return "../library/" + s.Id + ".m4a"
}

func download(s Song) {
	cmd := exec.Command("../bin/download", "-o", s.Id+".m4a", "-P", "../library", "-f", "m4a", s.Url)

	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb

	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(outb.String())
		fmt.Println(errb.String())
	}
}
