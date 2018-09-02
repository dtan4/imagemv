package image

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"
)

// Image represents image object
type Image struct {
	path string
}

// New creates image object
func New(path string) *Image {
	return &Image{
		path: path,
	}
}

// SHA1Sum returns SHA-1 hash of image
func (i *Image) SHA1Sum() (string, error) {
	body, err := ioutil.ReadFile(i.path)
	if err != nil {
		return "", errors.Wrapf(err, "cannot read image file %q", i.path)
	}

	return fmt.Sprintf("%x", sha1.Sum(body)), nil
}
