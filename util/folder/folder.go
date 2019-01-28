package folder

import "os"

// FolderIsExist
// 判断文件夹是否存在
func FolderExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// IsDir
// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IsFile
// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}
