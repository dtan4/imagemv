package fileutil

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"
)

// SHA1Sum returns SHA-1 hash of file
func SHA1Sum(path string) (string, error) {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		return "", errors.Wrapf(err, "cannot read image file %q", path)
	}

	return fmt.Sprintf("%x", sha1.Sum(body)), nil
}
