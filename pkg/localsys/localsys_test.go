package localsys

import (
	"os"
	"testing"
)

func TestValidPath(t *testing.T) {
	// Might be tricky but root is always a present path.
	// Might also encounter validation errors due to missing privileges.
	localPath := "/"
	if err := validatePath(localPath); err != nil {
		t.Errorf("Path validation failed: %s", err)
	}
}

func TestPathIsEmpty(t *testing.T) {
	var localPath string
	if err := validatePath(localPath); err != errPathIsEmpty {
		t.Errorf("Expected error to be %s, got %s", errPathIsEmpty, err)
	}
}

func TestPathNotAbs(t *testing.T) {
	localPath := "test"
	if err := validatePath(localPath); err != errPathIsNotAbs {
		t.Errorf("Expected error to be %s, got %s", errPathIsNotAbs, err)
	}
}

func TestPathInexistent(t *testing.T) {
	localPath := "/this/should/fail"
	if err := validatePath(localPath); !os.IsNotExist(err) {
		t.Errorf("Expected os.IsNotExistError, got %s", err)
	}
}

func TestValidateExt(t *testing.T) {
	fileName := "this.is.a.simple.test.jpeg"
	if !validateExt(fileName) {
		t.Error("Extension validation did not pass!")
	}
}

func TestInvalidExt(t *testing.T) {
	fileName := "this.is.another.test.go"
	if validateExt(fileName) {
		t.Error("Extension validation should not pass!")
	}
}
