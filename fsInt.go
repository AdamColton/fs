package fs

import (
	"io"
	"os"
)

var Std O = StdFS{}

type Filesystem interface {
	Open(name string) (File, error)
	Stat(name string) (os.FileInfo, error)
	Create(name string) (File, error)
	Mkdir(name string, perm os.FileMode) error
	RemoveAll(path string) error
	Rename(oldpath, newpath string) error
}

type File interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Writer
	Stat() (os.FileInfo, error)
	Sync() error
}

// osFS implements fileSystem using the local disk.
type StdFS struct{}

func (StdFS) Open(name string) (file, error)            { return os.Open(name) }
func (StdFS) Stat(name string) (os.FileInfo, error)     { return os.Stat(name) }
func (StdFS) Create(name string) (file, error)          { return os.Create(name) }
func (StdFS) Mkdir(name string, perm os.FileMode) error { return os.Mkdir(name, perm) }
func (StdFS) RemoveAll(path string) error               { return os.RemoveAll(path) }
func (StdFS) Rename(oldpath, newpath string) error      { return os.Rename(oldpath, newpath) }
