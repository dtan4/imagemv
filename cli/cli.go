package cli

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"

	"github.com/dtan4/imagemv/image"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

// CLI represents CLI object
type CLI struct {
	stdout io.Writer
	stderr io.Writer
}

// New creates CLI object
func New(stdout, stderr io.Writer) *CLI {
	return &CLI{
		stdout: stdout,
		stderr: stderr,
	}
}

// Run executes main command logic
func (cli *CLI) Run(args []string) error {
	if len(args) < 1 {
		return errors.New("dir must be given")
	}
	dir := args[0]

	var m sync.Mutex
	eg, _ := errgroup.WithContext(context.Background())

	f := bufio.NewWriter(cli.stdout)
	defer f.Flush()

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		eg.Go(func() error {
			if info.IsDir() {
				return nil
			}

			i := image.New(path)

			sha1sum, err := i.SHA1Sum()
			if err != nil {
				return errors.Wrapf(err, "cannot calculate SHA-1 checksum of %q", path)
			}

			m.Lock()
			f.WriteString(fmt.Sprintf("%s\t%s\n", path, sha1sum))
			m.Unlock()

			return nil
		})

		if err := eg.Wait(); err != nil {
			return errors.Wrap(err, "something wrong occured during for some files")
		}

		return nil
	})
	if err != nil {
		return errors.Wrapf(err, "something wrong occured during walking dir %q", dir)
	}

	return nil
}
