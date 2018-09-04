package fileutil

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	"gopkg.in/h2non/filetype.v1"
)

// IsImage returns whether the file is image or not
func IsImage(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, errors.Wrapf(err, "cannot read image file %q", path)
	}
	defer f.Close()

	// filetype requires only the first 261 bytes to judge file type
	head := make([]byte, 261)
	f.Read(head)

	return filetype.IsImage(head), nil
}

// SHA1Sum returns SHA-1 hash of file
func SHA1Sum(path string) (string, error) {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		return "", errors.Wrapf(err, "cannot read image file %q", path)
	}

	return fmt.Sprintf("%x", sha1.Sum(body)), nil
}
