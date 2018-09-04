package cli

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
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

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		fmt.Fprintln(cli.stdout, path)

		return nil
	})
	if err != nil {
		return errors.Wrapf(err, "something wrong occured during walking dir %q", dir)
	}

	return nil
}
