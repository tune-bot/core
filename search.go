package core

import (
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
			if i >= numResults {
				break
			}

			song := Song{
				Title:  track.Title,
				Artist: getArtists(track.Artists),
				Album:  track.Album.Name,
				Code:   track.VideoID,
			}

			// Get album info
			album := findAlbum(track.Album.Name, track.Album.ID)
			if album != nil {
				year, _ := strconv.ParseUint(album.Year, 10, 16)
				song.Year = uint16(year)
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
