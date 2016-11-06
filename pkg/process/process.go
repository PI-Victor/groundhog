package process

import (
	"github.com/PI-Victor/groundhog/pkg/process"
	"log"
)

func processFile(path string) error {
	metaData, err := process.GetExifMetadata(path)
	if err != nil {
		return err
	}
	log.Println(metaData.dateTime)
}
