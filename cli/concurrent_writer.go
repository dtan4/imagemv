package cli

import (
	"bufio"
	"io"
	"sync"
)

type concurrentWriter struct {
	w *bufio.Writer
	sync.Mutex
}

func newConcurrentWriter(w io.Writer) *concurrentWriter {
	return &concurrentWriter{
		w: bufio.NewWriter(w),
	}
}

func (c *concurrentWriter) Flush() error {
	return c.w.Flush()
}

func (c *concurrentWriter) WriteString(s string) (int, error) {
	c.Lock()
	defer c.Unlock()

	return c.w.WriteString(s)
}
