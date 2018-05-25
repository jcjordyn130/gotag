// flac implements FLAC support for gotag.
package flac

import (
	"strings"
	"strconv"
	"errors"
	"fmt"
	"github.com/mewkiz/flac"
	"github.com/mewkiz/flac/meta"
	"github.com/jcjordyn130/gotag/types"
	"github.com/jcjordyn130/gotag/decoders"
)

// metadata is the struct that contains the metadata and the FLAC stream.
type metadata struct {
	// file stores the currently open filename.
	file string

	// stream stores the open FLAC stream.
	stream *flac.Stream
}

// tagmap contains a list of vorbis tag names to try, as they aren't offically standardized.
var tagmap = map[string][]string {
	"track": []string{"TRACK", "TRACKNUMBER"},
	"tracktotal": []string{"TRACKTOTAL", "TOTALTRACKS"},
	"disc": []string{"DISCNUMBER"},
	"disctotal": []string{"TOTALDISCS"},
	"artist": []string{"ARTIST"},
	"composer": []string{"COMPOSER"},
	"album": []string{"ALBUM"},
	"title": []string{"TITLE"},
	"genre": []string{"GENRE", "STYLE"},
	"date": []string{"DATE"},
}

// getTag() returns the given tag.
// It doesn't actually use the tag it's provided, it uses the values from the tagmap.
func (m *metadata) getTag(tag string) ([]string, error) {
	// This is the variable that stores multible tag values.
	values := []string{}

	// Process all the FLAC blocks.
	for _, block := range m.stream.Blocks {
		// Try to get the specialised interface (type casting), if we can't then it's not a VorbisComment.
		body, ok := block.Body.(*meta.VorbisComment)
		if ! ok {
			continue
		}

		// Process the tags in the VorbisComment.
		for _, btag := range body.Tags {
			// Skip tags that are just a blank string.
			if btag[1] == "" {
				continue
			}

			// Process all the tags in tagmap.
			for _, mtag := range tagmap[tag] {
				// We use EqualFold, as tags are case sensitive and they aren't standardized.
				if strings.EqualFold(btag[0], mtag) {
					values = append(values, btag[1])
				}
			}
		}
	}

	// Return the tags, and an error if the values array is blank.
	if len(values) > 0 {
		return values, nil
	} else {
		return values, fmt.Errorf("Tag '%s' doesn't exist", tag)
	}
}

// AlbumArt() returns the embedded album art of a track.
// To make the code simpler we don't support more than one PictureBlock.
func (m *metadata) AlbumArt() (*types.Picture, error) {
	// Make a new Picture.
	picture := new(types.Picture)

	// Process all the FLAC blocks.
	for _, block := range m.stream.Blocks {
		// Try to get the specialised interface, if we can't then it's not a PictureBlock.
		body, ok := block.Body.(*meta.Picture)
		if ! ok {
			continue
		}

		// Fill in the Picture.
		picture.Data = body.Data
		picture.Height = int(body.Height)
		picture.Width = int(body.Width)
		picture.Mime = body.MIME

		// Return the picture.
		return picture, nil
	}

	// If we're here we couldn't get a picture.
	return nil, errors.New("No picture found")
}

// Sum() returns the sum of a track.
func (m *metadata) Sum() (string, error) {
	return decoders.Sum(m)
}

// File() returns the file path that a track represents.
func (m *metadata) File() (string) {
	return m.file
}

// Artist() returns the artist(s) of a track.
func (m *metadata) Artist() ([]string, error) {
	return m.getTag("artist")
}

// Composer() returns the composer(s) of a track.
func (m *metadata) Composer() ([]string, error) {
	return m.getTag("composer")
}

// Title() returns the title of a track.
func (m *metadata) Title() (string, error) {
	// Get the title.
	tag, err := m.getTag("title")
	if err != nil {	
		return "", err
	} else {
		// Return the album and any error.
		return tag[0], err
	}
}

// Album() returns the album(s) of a track.
func (m *metadata) Album() (string, error) {
	// Get the album.
	album, err := m.getTag("album")
	if err != nil {
		return "", err
	} else {
		// Return the album and any error.
		return album[0], nil
	}
}

// Year() returns the year a track was made.
func (m *metadata) Year() (int, error) {
	// Get the year.
	tag, err := m.getTag("date")
	if err != nil {
		return 0, err
	}

	// Convert it to an integer and return it.
	return strconv.Atoi(tag[0])
}

// Track() returns the track number.
func (m *metadata) Track() (int, error) {
	// Get the track number.
	tracknum, err := m.getTag("track")
	if err != nil {
		return 0, nil
	}

	// Convert it to an int.
	tracknum_int, err := strconv.Atoi(tracknum[0])
	if err != nil {
		return 0, err
	}

	// Return it.
	return tracknum_int, nil
}

// TotalTracks() returns the total number of tracks.
func (m *metadata) TotalTracks() (int, error) {
	// Get the number of total tracks.
	tracktotal, err := m.getTag("tracktotal")
	if err != nil {
		return 0, err
	}

	// Convert it to an int.
	tracktotal_int, err := strconv.Atoi(tracktotal[0])
	if err != nil {
		return 0, err
	}

	// Return it.
	return tracktotal_int, err
}

// Disc() returns the disc number.
func (m *metadata) Disc() (int, error) {
	// Get the disc number.
	disc, err := m.getTag("disc")
	if err != nil {
		return 0, err
	}

	// Convert it to an int.
	disc_int, err := strconv.Atoi(disc[0])
	if err != nil {
		return 0, err
	}

	// Return it.
	return disc_int, nil
}

// TotalDiscs() returns the number of discs in an album.
func (m *metadata) TotalDiscs() (int, error) {
	// Get the total disc number.
	disctotal, err := m.getTag("disctotal")
	if err != nil {
		return 0, err
	}

	// Convert it to an int.
	disctotal_int, err := strconv.Atoi(disctotal[0])
	if err != nil {
		return 0, err
	}

	// Return it.
	return disctotal_int, nil
}

// Genre() returns the genre(s) of a track.
func (m *metadata) Genre() ([]string, error) {
	return m.getTag("genre")
}

// Open() returns an open track.
func Open(file string) (*metadata, error) {
	// Define the variable we use to store errors.
	var err error

	// Make a new metadata struct.
	metadata := new(metadata)

	// Open the file.
	metadata.stream, err = flac.ParseFile(file)
	if err != nil {
		return nil, err
	}

	// Stuff the filename in the struct.
	metadata.file = file

	// Return the metadata object.
	return metadata, nil
}
