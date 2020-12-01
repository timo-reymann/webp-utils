package util

import (
	"errors"
	"path/filepath"
)

// FileGlobMatches returns all files matching the file util flag
func FileGlobMatches(pattern string) ([]string, error) {
	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, errors.New("glob doesnt match any file")
	}

	return files, nil
}
