// ogg implements OGG (vorbis/flac/opus/etc) support gotag.
package ogg

import (
	"github.com/xlab/vorbis-go/decoder"
	"github.com/jcjordyn130/gotag/types"
	//"github.com/jcjordyn130/gotag/decoders"
)

// metadata is the struct that contains the metadata and the OGG stream.
type metadata struct {
	// file stores the currently open filename.
	file string

	// stream stores the open OGG stream.
	stream *decoder.Decoder
}

func (m *metadata) AlbumArt() (*types.Picture, error) {
	return nil, errors.New("Not implemented")
}

func (m *metadata) Sum() (string, error) {
	return "", errors.New("Not implemented")
}

func (m *metadata) File() (string) {
	return ""
}

func (m *metadata) Artist() ([]string, error) {
	return []string{""}, errors.New("Not implemented")
}

func (m *metadata) Composer() ([]string, error) {
	return []string{""}, errors.New("Not implemented")
}

func (m *metadata) Title() (string, error) {
	return "", errors.New("Not implemented")
}

func (m *metadata) Album() (string, error) {
	return "", errors.New("Not implemented")
}

func (m *metadata) Year() (int, error) {
	return 0, errors.New("Not implemented")
}

func (m *metadata) Track() (int, error) {
	return 0, errors.New("Not implemented")
}

func (m *metadata) TotalTracks() (int, error) {
	return 0, errors.New("Not implemented")
}

func (m *metadata) Disc() (int, error) {
	return 0, errors.New("Not implemented")
}

func (m *metadata) TotalDiscs() (int, error) {
	return 0, errors.New("Not implemented")
}

func (m *metadata) Genre() ([]string, error) {
	return []string{""}, errors.New("Not implemented")
}

func (m *metadata) Open(file string) (*metadata, error) {
	return nil, errors.New("Not implemented yet")
	// Define the variable we use the store errors.
	//var err error

	// Make a new metadata struct.
	//metadata := new(metadata)

	// Open the file.
	//openfile, err := os.Open(file)
	//if err != nil {
	//	return nil, err
	//}

	// Open the decoder.
	//metadata.stream, err = decoder.New
}
