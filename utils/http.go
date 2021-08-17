package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

func RequestGet(url string) []byte {
	//resp, err := http.Get(url)
	//if err != nil {
	//	glog.Errorf("url错误1: %s", url)
	//	return nil
	//}

	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//// err != nil
	//if resp.StatusCode != 200 {
	//	glog.Errorf("url错误2: %s", url)
	//	return nil
	//}
	//return body
	client := CreateClient()
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("httpGet url : %v , error : %v\n", url, err)
		return nil
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("httpGet url : %v , error : %v\n", url, err)
		}
	}()
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("httpGet url : %v , error : %v\n", url, err)
		return nil
	}
	//return string(body)
	return body
}

// Fetch from url
func Fetch(url string) io.Reader {
	client := CreateClient()
	resp, err := client.Get(url)
	if err != nil {
		log.Fatalf("Fetch url from %s error: %s", url, err)
	}
	if resp.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", resp.StatusCode, resp.Status)
	}
	return resp.Body
}
