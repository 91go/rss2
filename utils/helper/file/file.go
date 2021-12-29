package file

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gogf/gf/os/gfile"
)

// 获取文件MimeType
func GetContentType(filePath string) string {
	file, err := gfile.Open(filePath)
	if err != nil {
		return "audio/mpeg"
	}
	buffer := make([]byte, 512)

	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println(err)
	}
	contentType := http.DetectContentType(buffer)
	return contentType
}

func GetAllFiles(dir string) ([]string, error) {
	dirPath, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var files []string
	sep := string(os.PathSeparator)
	for _, fi := range dirPath {
		if fi.IsDir() { // 如果还是一个目录，则递归去遍历
			subFiles, err := GetAllFiles(dir + sep + fi.Name())
			if err != nil {
				return nil, err
			}
			files = append(files, subFiles...)
		} else {
			files = append(files, dir+sep+fi.Name())
		}
	}
	return files, nil
}
