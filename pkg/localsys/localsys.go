package localsys

import (
	"errors"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const (
	jpeg = "jpeg"
	jpg  = "jpg"
	bmp  = "bmp"
	png  = "png"
)

var (
	// okFormats are the accepted formats for the files that we want to rename.
	okFormats = []string{jpeg, jpg, bmp, png}

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
	if err = controller(dir); err != nil {
		log.Fatal(err)
	}
}

// controller dispatches a new delegator in a goroutine for each directory/file
// it encounters.
func controller(dir string) error {
	// i want a channel that returns errors to the controller, up the stack, from
	// the spawned delegators. a new delegator is spawned only if the file is a
	// directory.

	//errOutput := make(chan error)
	go filepath.Walk(dir, delegator)
	return nil
}

// delegator is called for each file in the specified directory and analyses if
// the file is a simple file or a directory. If it is a directory it will call
// itself recursively and spawn a new delegator.
func delegator(path string, info os.FileInfo, err error) error {
	fileMode := info.Mode()
	if fileMode.IsDir() {
		go filepath.Walk(path, delegator)
	}

	if fileMode.IsRegular() && validateExt(info.Name()) {

		return nil
	}
	return nil
}

// validateExt checks if the file contains a valid extention.
func validateExt(fileName string) bool {
	fileExt := fileName[strings.LastIndex(fileName, ".")+1:]
	for _, format := range okFormats {
		if fileExt == format {
			return true
		}
	}
	return false
}
