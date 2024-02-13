package entities

import (
	"fmt"
	"time"
)

type Album struct {
	Title       string
	Artist      string
	ReleaseDate time.Time
}

func NewAlbum(title string, artist string, releaseDate time.Time) *Album {
	return &Album{
		Title:       title,
		Artist:      artist,
		ReleaseDate: releaseDate,
	}
}

func (a *Album) String() string {
	return fmt.Sprintf("Album Title: %s, by %s, released on %s", a.Title, a.Artist, a.ReleaseDate.Format("Jan 2, 2006"))
}
