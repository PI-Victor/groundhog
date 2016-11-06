package exif

import (
	"github.com/rwcarlsen/goexif/exif"
	"os"
	"time"
)

// MediaFile holds the metadata of a image file.
type MediaFile struct {
	name     string
	dateTime time.Time
}

func newMediaFile(fileName string, metaData *exif.Exif) (*MediaFile, error) {
	date, err := metaData.DateTime()
	if err != nil {
		return nil, err
	}
	return &MediaFile{
		name:     fileName,
		dateTime: date,
	}, nil
}

// GetExifMetadata returns a new MediaFile struct with the metadata of the
// image file.
func GetExifMetadata(path string) (*MediaFile, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	metaData, err := exif.Decode(file)
	if err != nil {
		return nil, err
	}
	return newMediaFile("test", metaData)
}
