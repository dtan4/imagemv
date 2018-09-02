package image

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	path := "/path/to/image"
	image := New(path)

	if image == nil {
		t.Errorf("image must be created")
	}
}

func TestSha1Sum(t *testing.T) {
	testcases := []struct {
		path string
		want string
	}{
		{
			path: filepath.Join("..", "testdata", "320x200.jpg"),
			want: "999877f1981a45f9080534fcb9a3e3371269010b",
		},
		{
			path: filepath.Join("..", "testdata", "320x200.png"),
			want: "1131c73fbff8894833d2710bafc66ba008879e3a",
		},
	}

	for _, tc := range testcases {
		image := &Image{
			path: tc.path,
		}

		got, err := image.SHA1Sum()
		if err != nil {
			t.Errorf("got error: %s", err)
		}

		if got != tc.want {
			t.Errorf("want: %q, got: %q", tc.want, got)
		}
	}
}

func TestSha1Sum_error(t *testing.T) {
	testcases := []struct {
		path string
		msg  string
	}{
		{
			path: "filedoesnotexist",
			msg:  `cannot read image file "filedoesnotexist"`,
		},
	}

	for _, tc := range testcases {
		image := &Image{
			path: tc.path,
		}

		_, err := image.SHA1Sum()
		if err == nil {
			t.Errorf("no error")
		}

		if !strings.Contains(err.Error(), tc.msg) {
			t.Errorf("want in error: %q, got: %q", tc.msg, err.Error())
		}
	}
}
