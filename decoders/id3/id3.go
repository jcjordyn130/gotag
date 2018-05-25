// id3 implements ID3 (mp3, etc) support for gotag.
package id3

import (
	"github.com/bogem/id3v2"
	"github.com/jcjordyn130/gotag/types"
	"github.com/jcjordyn130/gotag/decoders"
	"errors"
	"strconv"
	"strings"
	"fmt"
)

// metadata is the struct that contains the metadata and the ID3 stream.
type metadata struct {
	// file stores the currently open filename.
	file string

	// stream stores the open ID3 stream.
	stream *id3v2.Tag
}

// getTag() returns the given tag, or an error if it doesn't exist.
func (m *metadata) getTag(tag string) (string, error) {
	stag := m.stream.GetTextFrame(tag).Text

	if stag != "" {
		return stag, nil
	} else {
		return stag, errors.New("Tag doesn't exist")
	}		
}

// getTagInt() returns the given tag as an integer.
func (m *metadata) getTagInt(tag string) (int, error) {
	// Get the tag.
	stag, err := m.getTag(tag)
	if err != nil {
		return 0, err
	}

	// If it's blank, return 0.
	if stag == "" {
		return 0, errors.New("Tag doesn't exist")
	}

	// Remove the null bytes.
	stag = strings.Trim(stag, "\x00")

	// Convert it to an integer.
	stag_int, err := strconv.Atoi(stag)
	if err != nil {
		return 0, err
	}

	// Return it.
	return stag_int, nil
}

// getTags() returns the given tag as a string array.
// We need this as ID3 doesn't support more than one text frame for a tag, so delimiter hell time.
func (m *metadata) getTags(tag string) ([]string, error) {
	// Get the tag.
	stag, err := m.getTag(tag)
	if err != nil {
		return []string{}, err
	}

	// Process the delimiter
	if m.stream.Version() == 3 {
		// We're using ID3 v2.3, so use /, :, or ;.
		// Split the tag by /.
		tags := deleteBlank(strings.Split(stag, "/"))

		// Split the tag by :.
		tagc := deleteBlank(strings.Split(stag, ":"))

		// Split the tag by ;.
		tagsc := deleteBlank(strings.Split(stag, ";"))

		// Return any of the split tags that have a length over 1.
		if len(tags) > 1 {
			return tags, nil
		}

		if len(tagc) > 1 {
			return tagc, nil 
		}

		if len(tagsc) > 1 {
			return tagsc, nil
		}

		// If we're here it means the tag doesn't have a delimiter, so just return it.
		return []string{stag}, nil
	} else if m.stream.Version() == 4 {
		// We're using ID3 v2.4, so use a null char.
		taga := strings.Split(stag, "\x00")
		return taga, nil
	} else {
		// We're using an unknown version.
		return []string{}, fmt.Errorf("Unknown ID3 version: %d \n", m.stream.Version)
	}
}

// deleteBlank removes blank strings from a string array.
func deleteBlank(s []string) []string {
	// Make a new array.
	var r []string

	// Loop though the old array, adding non-blank strings to the new array.
	for _, str := range s {
		if str != "" || str != " " {
			r = append(r, str)
		}
	}

	// Return the new array.
	return r
}

// AlbumArt() returns the embedded album art of a track.
// To make the code simpler we don't support more than one PictureBlock.
func (m *metadata) AlbumArt() (*types.Picture, error) {
	return nil, errors.New("Not implemented")
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
	return m.getTags(m.stream.CommonID("Artist"))
}

// Composer() returns the composer(s) of a track.
func (m *metadata) Composer() ([]string, error) {
	return m.getTags(m.stream.CommonID("Composer"))
}

// Title() returns the title of a track.
func (m *metadata) Title() (string, error) {
	return m.getTag(m.stream.CommonID("Title"))
}

// Album() returns the album(s) of a track.
func (m *metadata) Album() (string, error) {
	return m.getTag(m.stream.CommonID("Album/Movie/Show title"))
}

// Year() returns the year a track was made.
func (m *metadata) Year() (int, error) {
	return m.getTagInt(m.stream.CommonID("Year"))
}

// Track() returns the current track.
func (m *metadata) Track() (int, error) {
	// Get the track number.
	track, err := m.getTags(m.stream.CommonID("Track number/Position in set"))
	if err != nil {
		return 0, err
	}

	// If it's a blank, return 0.
	if track[0] == "" {
		return 0, errors.New("Tag doesn't exist")
	}

	// Convert it to an integer.
	// HACK: we assume the format is currenttrack delim totaltracks
	track_int, err := strconv.Atoi(track[0])
	if err != nil {
		return 0, err
	}

	// Return it.
	return track_int, nil
}

// Track() returns the total number of track.s
func (m *metadata) TotalTracks() (int, error) {
	// Get the track number.
	tracktotal, err := m.getTags(m.stream.CommonID("Track number/Position in set"))
	if err != nil {
		return 0, err
	}

        // If it only has the current track, then error out.
	if tracktotal[1] == "" {
		return 0, errors.New("Tag doesn't exist")
	}

	// Convert it to an integer.
	// HACK: we assume the format is currenttrack delim totaltracks
	tracktotal_int, err := strconv.Atoi(tracktotal[1])
	if err != nil {
		return 0, err
	}

	// Return it.
	return tracktotal_int, nil
}

// Disc() returns the disc number.
func (m *metadata) Disc() (int, error) {
	// Get the disc number.
	disc, err := m.getTags(m.stream.CommonID("Part of a set"))
	if err != nil {
		return 0, err
	}

        // If it's a blank, return 0.
	if len(disc) == 0 {
		return 0, errors.New("Tag doesn't exist")
	}

	// Convert it to an integer.
	// HACK: we assume the format is currentdisc delim totaldiscs
	disc_int, err := strconv.Atoi(disc[0])
	if err != nil {
		return 0, err
	}

	// Return it.
	return disc_int, nil
}

// TotalDiscs() returns the number of discs in an album.
func (m *metadata) TotalDiscs() (int, error) {
	// Get the number of total discs.
	disctotal, err := m.getTags(m.stream.CommonID("Part of a set"))
	if err != nil {
		return 0, err
	}

        // If it's a blank, return 0.
	if len(disctotal) == 0 {
		return 0, errors.New("Tag doesn't exist")
	}

	// Convert it to an integer.
	// HACK: we assume the format is currentdisc delim totaldiscs
	disctotal_int, err := strconv.Atoi(disctotal[1])
	if err != nil {
		return 0, err
	}

	// Return it.
	return disctotal_int, nil
}

// Genre() returns the genre(s) of a track.
func (m *metadata) Genre() ([]string, error) {
	// Get the tag.
	tag, err := m.getTag(m.stream.CommonID("Content type"))

	// I think ID3 only supports one genre type, so return it in an array for compatibilty.
	return []string{tag}, err
}

// Open() returns an open track.
func Open(file string) (*metadata, error) {
	// Define the variable we use to store errors.
	var err error

	// Make a new metadata struct.
	metadata := new(metadata)

	// Open the file.
	metadata.stream, err = id3v2.Open(file, id3v2.Options{Parse: true})
	if err != nil {
		return nil, err
	}

	// Stuff the filename in the struct
	metadata.file = file

	// Return the metadata object.
	return metadata, nil
}
