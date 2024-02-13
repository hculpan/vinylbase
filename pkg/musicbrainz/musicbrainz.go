package musicbrainz

import (
	"fmt"
	"log"

	"github.com/michiwend/gomusicbrainz"
)

func fetchAlbum(artistName string, title string) {
	client, err := gomusicbrainz.NewWS2Client(
		"https://musicbrainz.org/ws/2",
		"VinylBase",
		"0.0.1-beta",
		"http://github.com/hculpan/vinylbase")
	if err != nil {
		log.Fatal(err)
	}

	// Search for some artist(s)
	resp, _ := client.SearchReleaseGroup(fmt.Sprintf(`release-group:%q AND artist:%q`, artistName, title), -1, -1)

	// Pretty print Name and score of each returned artist.
	for _, rg := range resp.ReleaseGroups {
		fmt.Printf("Name: %-50s Release Date: %-25s\n", rg.Title, rg.FirstReleaseDate.Format("2006-01-02"))
	}

}
