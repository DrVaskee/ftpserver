package server

import (
	"os"
	"time"
)

type VirtualFileInfo struct {
	Internal_name string
	Internal_size int64
	Internal_mode os.FileMode
}

func (f VirtualFileInfo) Name() string {
	return f.Internal_name
}

func (f VirtualFileInfo) Size() int64 {
	return f.Internal_size
}

func (f VirtualFileInfo) Mode() os.FileMode {
	return f.Internal_mode
}

func (f VirtualFileInfo) IsDir() bool {
	return f.Internal_mode.IsDir()
}

func (f VirtualFileInfo) ModTime() time.Time {
	return time.Now().UTC()
}

func (f VirtualFileInfo) Sys() interface{} {
	return nil
}
