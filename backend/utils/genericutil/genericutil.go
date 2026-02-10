package genericutil

import (
	"os"
)

const USER_ONLY_FILE_PERMS = 0600
const USER_ONLY_DIR_PERMS = 0700

// NarrowCreate creates a file with narrow access permissions (limited to the current user).
// NOTE: Does not close file with defer! Caller should call as:
//
// f, err := NarrowCreate(filepath)
// if err != nil { /* error handling / escape */ }
// defer f.Close()

func NarrowCreate(fpath string) (*os.File, error) {
	file, err := os.Create(fpath)
	if err != nil {
		return nil, err
	}
	err = file.Chmod(USER_ONLY_FILE_PERMS)
	if err != nil {
		closeErr := file.Close()
		if closeErr != nil {
			return nil, closeErr
		}
		return nil, err
	}
	return file, nil
}

