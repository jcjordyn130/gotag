package gotag

import (
	"fmt"
	"io/ioutil"
	"errors"
	"gopkg.in/h2non/filetype.v1"
	"github.com/jcjordyn130/gotag/decoders/flac"
	"github.com/jcjordyn130/gotag/decoders/id3"
	"github.com/jcjordyn130/gotag/types"
)

func getMime(file string) (string, error) {
	buf, _ := ioutil.ReadFile(file)

	kind, unknown := filetype.Match(buf)
	if unknown != nil {
		return "", errors.New("Unknown MIME type")
	}

	return kind.MIME.Value, nil
}

func Open(file string) (types.Metadata, error) {
	// Get the mime type.
	mime, err := getMime(file)
	if err != nil {
		return nil, err
	}

	// Open the right decoder for the file, if we don't have one return an error.
	// If you know what file you have before hand you can use the decoders manually.
	if mime == "audio/x-flac" {
		return flac.Open(file)
	} else if mime == "audio/mpeg" {
		return id3.Open(file)
	} else {
		return nil, fmt.Errorf("Unsupported audio format: " + mime)
	}
}
