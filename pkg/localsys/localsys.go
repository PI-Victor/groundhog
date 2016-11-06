package localsys

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
)

const (
	jpeg = "jpeg"
	jpg  = "jpg"
	bmp  = "bmp"
	png  = "png"
)

var (
	// okFormats are the accepted formats for the files that we want to rename.
	okFormats       = []string{jpeg, jpg, bmp, png}
	errPathIsEmpty  = errors.New("Path is empty!")
	errPathIsNotAbs = errors.New("Path is not absolute!")
)

func validatePath(localPath string) error {
	if localPath == "" {
		return errPathIsEmpty
	}
	if !path.IsAbs(localPath) {
		return errPathIsNotAbs
	}
	if _, err := os.Stat(localPath); err != nil {
		return err
	}

	return nil
}

// Run searches the path for files.
func Run(localPath string) {
	if err := validatePath(localPath); err != nil {
		log.Fatal(err)
	}
	dir, err := filepath.Abs(localPath)
	if err != nil {
		log.Fatalf("An error occured while reading the %s directory: %s", localPath, err)
	}

	if err := filepath.Walk(dir, delegator); err != nil {
		log.Fatal(err)
	}
}

// delegator is called for each file in the specified directory and analyses if
// the file is a simple file or a directory. If it is a directory it will call
// itself recursively and spawn a new delegator.
func delegator(path string, info os.FileInfo, err error) error {
	fileMode := info.Mode()
	outErr := make(chan error)
	if fileMode.IsDir() {
		go func() {
			outErr <- filepath.Walk(path, delegator)
		}()
		return fmt.Errorf("%v", outErr)
	}
	if fileMode.IsRegular() && validateExt(info.Name()) {
		return nil
	}
	return nil
}

// validateExt checks if the image file contains a valid extention.
func validateExt(fileName string) bool {
	for _, format := range okFormats {
		match, err := filepath.Match("*."+format, fileName)
		if err != nil || !match {
			continue
		}
		return true
	}
	return false
}
