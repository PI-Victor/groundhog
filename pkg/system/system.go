package localsys

import (
	"errors"
	"log"
	"os"
	"path"
	"path/filepath"
)

func validatePath(localPath string) error {
	if localPath == "" {
		return errors.New("Path sent as argument is empty!")
	}
	if !path.IsAbs(localPath) {
		return errors.New("Path is not absolute!")
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
	filepath.Walk(dir, printDir)
}

func printDir(path string, info os.FileInfo, err error) error {
	log.Println(path)
	log.Println(info)
	return nil
}

// makeBackup creates a backup of the file, before renaming it and storing it
// in the directory specified.
func makeBackup() error {
	return nil
}
