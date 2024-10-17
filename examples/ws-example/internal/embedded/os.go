package embedded

import (
	"io/fs"
	"os"
)

type OsFs struct {
}

func (receiver OsFs) Open(name string) (fs.File, error) {
	return os.Open(name)
}

func NewOsFs() OsFs {
	return OsFs{}
}
