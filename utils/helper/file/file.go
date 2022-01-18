package file

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gogf/gf/net/ghttp"

	"github.com/gogf/gf/os/gfile"
)

const (
	DefaultMimeType = "audio/mpeg"
)

// 获取文件MimeType
func GetLocalFileMimeType(filePath string) string {
	file, err := gfile.Open(filePath)
	if err != nil {
		return DefaultMimeType
	}
	defer file.Close()

	buffer := make([]byte, 512)

	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return DefaultMimeType
	}
	contentType := http.DetectContentType(buffer)
	return contentType
}

// 获取远程文件的MIME Type
func GetRemoteFileMIMEType(url, username, pwd string) string {
	resp, err := ghttp.NewClient().SetBasicAuth(username, pwd).Get(url)
	if err != nil {
		return DefaultMimeType
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return DefaultMimeType
	}
	return http.DetectContentType(bytes)
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
