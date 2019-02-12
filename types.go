package fs

import (
	"io"
	"os"
	"time"
)

// File interface describes a generic file methods.
type File interface {
	Object

	io.Reader
	io.ReaderAt
	io.ReaderFrom
	io.WriteSeeker
	io.WriterAt
	io.WriterTo
	io.Closer
}

type Dir interface {
	Object

	CreateFile(path string, perm os.FileMode, body ...[]byte) (f File, err error)
	CreateDir(path string, perm os.FileMode) (d Dir, err error)
	Read() ([]Object, error)
	ReadFiles() ([]File, error)
}

type Object interface {
	CreatedAt() time.Time
	UpdatedAt() time.Time
	Readable() (bool, error)
	Writeable() (bool, error)

	IsFile() (f File, ok bool)
	IsDir() (d Dir, ok bool)
}

type Basic interface {
	DirExists(path string) (exists bool, err error)
	ReadFile(path string) (body []byte, err error)
	OpenFile(path string) (f File, err error)

	FileExists(path string) (exists bool, err error)
	ReadFilesInDir(path string) (files []File, err error)
	OpenDir(path string) (d Dir)

	Chroot(path string) Basic
}

type Chmoder interface {
	Chmod(path string, perm os.FileMode) (changed bool, err error)
}

type Chowner interface {
	Chown(path string, uid interface{}, gid interface{}) (changed bool, err error)
}

type Unchrooter interface {
	Unchroot(path string)
}

type Permissions interface {
	Permissions() os.FileMode
}
