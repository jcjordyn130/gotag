// gettag uses gotag to output the tags of a file.
package main

import (
	"github.com/jcjordyn130/gotag"
	"fmt"
	"os"
	"path/filepath"
	//"runtime/debug"
)


// handleError() handles errors, it just outputs it in a fancy way.
func handleError(file string, err error) {
	if err != nil {
		// Print out a stack trace.
		//debug.PrintStack()

		// Print out the error.
		fmt.Printf("[%s] Error: %s \n", filepath.Base(file), err)
	}
}

func main() {
	for num, arg := range os.Args {
		// Skip handling ourselfs.
		if num == 0 {
			continue
		}

		// Open the file.
		m, err := gotag.Open(arg)
		if err != nil {
			handleError(arg, err)
			continue
		}

		// Get the artists.
		Artist, err := m.Artist()
		handleError(arg, err)

		// Get the composers.
		Composer, err := m.Composer()
		handleError(arg, err)

		// Get the title.
		Title, err := m.Title()
		handleError(arg, err)

		// Get the album.
		Album, err := m.Album()
		handleError(arg, err)

		// Get the year.
		Year, err := m.Year()
		handleError(arg, err)

		// Get the track.
		Track, err := m.Track()
		handleError(arg, err)

		// Get the total tracks.
		TrackTotal, err := m.TotalTracks()
		handleError(arg, err)

		// Get the disc.
		Disc, err := m.Disc()
		handleError(arg, err)

		// Get the total discs.
		DiscTotal, err := m.TotalDiscs()
		handleError(arg, err)

		// Get the genres.
		Genre, err := m.Genre()
		handleError(arg, err)

		// Get the file path.
		File := filepath.Base(m.File())

		// Get the sum.
		Sum, err := m.Sum()
		handleError(arg, err)

		// Print the artists.
		for _, artist := range Artist {
			fmt.Printf("[%s] Artist: %s \n", File, artist)
		}
	
		// Print the composers.
		for _, composer := range Composer {
			fmt.Printf("[%s] Composer: %s \n", File, composer)
		}

		// Print the genres.
		for _, genre := range Genre {
			fmt.Printf("[%s] Genre: %s \n", File, genre)
		}

		// Print the other details.
		fmt.Printf("[%s] Album: %s \n", File, Album)
		fmt.Printf("[%s] Title: %s \n", File, Title)
		fmt.Printf("[%s] Year: %d \n", File, Year)
		fmt.Printf("[%s] Track: %d \n", File, Track)
		fmt.Printf("[%s] Total Tracks: %d \n", File, TrackTotal)
		fmt.Printf("[%s] Disc: %d \n", File, Disc)
		fmt.Printf("[%s] Total Discs: %d \n", File, DiscTotal)
		fmt.Printf("[%s] Sum: %s \n", File, Sum)
	}
}
