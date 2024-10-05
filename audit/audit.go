package audit

import "io/fs"

type Auditable interface {
	LastAccessTime(fs.FileInfo) int64
}
