// types contains shared types
package types

// Picture defines the album art of an audio file.
type Picture struct {
	Mime   string
	Data   []byte
	Width  int
	Height int
}

// Metadata defines the metadata of an audio file, decoders have to implement this.
type Metadata interface {
	AlbumArt()    (*Picture, error)
	Sum()         (string, error)
	File()        (string)
	Artist()      ([]string, error)
	Composer()    ([]string, error)
	Title()       (string, error)
	Album()       (string, error)
	Year()        (int, error)
	Track()       (int, error)
	TotalTracks() (int, error)
	Disc()        (int, error)
	TotalDiscs()  (int, error)
	Genre()       ([]string, error)
}
