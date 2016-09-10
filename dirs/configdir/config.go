package configdir

import (
	"path/filepath"
)

// the various file paths of config folder
var (
	InitrdLocalIncludeDir string

	LocalIncludeDir string

	WritableLocalIncludeDir string
)

func init() {
	SetRootDir("")
}

func SetRootDir(rootdir string) {
	InitrdLocalIncludeDir = filepath.Join(rootdir, "writable_local-include")

	LocalIncludeDir = filepath.Join(rootdir, "local-include")

	WritableLocalIncludeDir = filepath.Join(rootdir, "writable_local-include")
}
