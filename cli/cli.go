package cli

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/dtan4/imagemv/fileutil"

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

	eg, ctx := errgroup.WithContext(context.Background())
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	cso := newConcurrentWriter(cli.stdout)
	defer cso.Flush()

	cse := newConcurrentWriter(cli.stderr)
	defer cse.Flush()

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				return nil
			default:
				if info.IsDir() {
					return nil
				}

				b, err := fileutil.IsImage(path)
				if err != nil {
					return errors.Wrapf(err, "cannot judge whether %q is an image or not", path)
				}

				if !b {
					cse.WriteString(fmt.Sprintf("warning: %q is not an image\n", path))
					return nil
				}

				sha1sum, err := fileutil.SHA1Sum(path)
				if err != nil {
					return errors.Wrapf(err, "cannot calculate SHA-1 checksum of %q", path)
				}

				cso.WriteString(fmt.Sprintf("%s\t%s\n", path, sha1sum))
			}

			return nil
		})

		return nil
	})

	if err != nil {
		return errors.Wrapf(err, "something wrong occured during walking dir %q", dir)
	}

	if err := eg.Wait(); err != nil {
		cancel()
		return errors.Wrap(err, "something wrong occured during for some files")
	}

	return nil
}
