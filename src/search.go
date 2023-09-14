package core

import (
	"database/sql"
	"strconv"

	"github.com/raitonoberu/ytmusic"
)

func Search(query string, numResults int) []Song {
	results := []Song{}
	resp := ytmusic.TrackSearch(query)
	result, _ := resp.Next()

	if numResults > 5 {
		numResults = 5
	} else if numResults < 1 {
		numResults = 1
	}

	if result != nil {
		for i, track := range result.Tracks {
			// Only return top # search results
			if i == numResults {
				break
			}

			song := Song{
				Title:  sql.NullString{String: track.Title, Valid: true},
				Artist: sql.NullString{String: getArtists(track.Artists), Valid: true},
				Album:  sql.NullString{String: track.Album.Name, Valid: true},
				Code:   sql.NullString{String: track.VideoID, Valid: true},
			}

			// Get album info
			album := findAlbum(track.Album.Name, track.Album.ID)
			if album != nil {
				year, _ := strconv.ParseUint(album.Year, 10, 16)
				song.Year = sql.NullInt16{Int16: int16(year), Valid: true}
			}

			// Add song to search results
			results = append(results, song)
		}
	}

	return results
}

func getArtists(artists []ytmusic.Artist) string {
	artistNames := ""

	for i, artist := range artists {
		artistNames += artist.Name

		if i < len(artists)-1 {
			artistNames += ", "
		}
	}

	return artistNames
}

func findAlbum(title, id string) *ytmusic.AlbumItem {
	resp := ytmusic.AlbumSearch(title)
	var result *ytmusic.SearchResult
	var err error

	// Find the matching album
	for err == nil {
		result, err = resp.Next()

		if result != nil {
			for _, album := range result.Albums {
				if album.BrowseID == id {
					return album
				}
			}
		}
	}

	return nil
}
