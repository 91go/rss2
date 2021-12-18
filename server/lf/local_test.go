package lf

import (
	"fmt"
	"github.com/gogf/gf/text/gstr"
	"net/http"
	"os"
	"testing"
)

func TestGetAllFiles(t *testing.T) {
	// files, err := GetAllFiles("/Users/lhgtqb7bll/Downloads/ys")
	// if err != nil {
	// 	return
	// }

	// dirNames, err := gfile.DirNames("/Users/lhgtqb7bll/Downloads/ys/Bitchery/年轻继母的尝试.mp3")
	// dirNames := gfile.Dir("/Users/lhgtqb7bll/Downloads/ys/Bitchery/年轻继母的尝试.mp3")
	//
	// fmt.Println(dirNames)

	str := gstr.Str("/srv/ys/23423434321/8.开门，你的小女仆请查收3快乐姐妹.mp3", "/srv")
	fmt.Println(str)

	// str := gstr.Str("/srv/ys/23423434321/8.开门，你的小女仆请查收3快乐姐妹.mp3", "/srv/")
	// fmt.Println(str)
}

func TestContentType(t *testing.T) {
	buffer := make([]byte, 512)
	file, err := os.Open("/Users/lhgtqb7bll/Documents/Xnip2021-11-30_20-00-57.jpg")

	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println(err)
	}
	contentType := http.DetectContentType(buffer)
	fmt.Println(contentType)
}
