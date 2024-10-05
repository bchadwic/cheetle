package audit

import (
	"io/fs"
	"syscall"
)

type Windows struct{}

func (w *Windows) LastAccessTime(info fs.FileInfo) int64 {
	stat := info.Sys().(*syscall.Win32FileAttributeData)
	return stat.LastAccessTime.Nanoseconds()
}
