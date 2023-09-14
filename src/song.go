package core

import (
	"bytes"
	"database/sql"
	"fmt"
	"os/exec"
	"strings"

	"github.com/google/uuid"
)

type Song struct {
	Id     sql.NullString `json:"id"`
	Code   sql.NullString `json:"code"`
	Title  sql.NullString `json:"title"`
	Artist sql.NullString `json:"artist"`
	Album  sql.NullString `json:"album"`
	Year   sql.NullInt16  `json:"year"`
}

func (s *Song) AddToPlaylist(playlistId string) error {
	songId := uuid.New().String()
	songExists := false
	_, err := db.Exec(`
		insert into song 
		(id, code, title, artist, album, year) 
		values 
		(uuid_to_bin(?), ?, ?, ?, ?, ?);`,
		songId, s.Code, s.Title, s.Artist, s.Album, s.Year)

	if err != nil {
		result, err := db.Query(`
			select bin_to_uuid(id) as id 
			from song 
			where code = ?;`,
			s.Code)

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

	s.Id = sql.NullString{String: songId, Valid: true}

	if !songExists {
		go s.download()
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
	return "../library/" + s.Id.String + ".m4a"
}

func (s Song) String() string {
	str := fmt.Sprintf("Title: %s\n", s.Title)

	if strings.Count(s.Artist.String, ",") > 0 {
		str += "Artists: "
	} else {
		str += "Artist: "
	}
	str += s.Artist.String + "\n"
	str += fmt.Sprintf("Album: %s\n", s.Album)

	if s.Year.Int16 != 0 {
		str += fmt.Sprintf("Year: %d\n", s.Year)
	}

	return str
}

func (s Song) download() {
	cmd := exec.Command("../bin/download", "-o", s.Id.String+".m4a", "-P", "../library", "-f", "m4a", fmt.Sprintf("https://music.youtube.com/watch?v=%s", s.Code))

	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb

	err := cmd.Run()
	if err != nil {
		PrintError(err.Error())
	} else {
		PrintSuccess(outb.String())
		PrintError(errb.String())
	}
}
