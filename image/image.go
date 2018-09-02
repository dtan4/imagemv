package image

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
