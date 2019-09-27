package common

import "os"

// JudgeType judge the type of the name.
func JudgeType(name string) (bool, error) {
	fileInfo, err := os.Stat(name)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), nil
}
