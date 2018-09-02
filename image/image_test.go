package image

import (
	"testing"
)

func TestNew(t *testing.T) {
	path := "/path/to/image"
	image := New(path)

	if image == nil {
		t.Errorf("image must be created")
	}
}
