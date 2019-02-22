/**
 * A few function in php
 */

package hls

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

//检查文件是否存在
func IsFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}

//检查目录是否存在
func IsDir(path string) bool {
	p, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	} else {
		return p.IsDir()
	}
}

//写内容到文件
func FilePutContents(filename string, content string) (val bool, err error) {
	if IsFile(filename) {
		return
	}

	fout, err := os.Create(string(filename))

	if err != nil {
		return false, err
	}

	defer fout.Close()

	fout.WriteString(content)

	return true, nil
}

func FileGetContents(uri string) (str string, err error) {
	//url := strings.ToLower(strings.TrimSpace(uri))
	url := uri

	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		client := &http.Client{}
		//提交请求
		request, err := http.NewRequest("GET", url, nil)

		//增加header选项
		//request.Header.Add("Cookie", "xxxxxx")
		request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36")
		//request.Header.Add("X-Requested-With", "xxxx")

		if err != nil {
			panic(err)
		}
		resp, err := client.Do(request)

		if err != nil {
			return "", err
		}

		defer resp.Body.Close()

		if resp.StatusCode >= 400 {
			return "", errors.New(resp.Status)
		}

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return "", err
		}

		return string(body), nil
	}

	data, err := ioutil.ReadFile(uri)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func Unlink(uri string) error {
	if IsFile(uri) {
		err := os.Remove(uri)
		return err
	}

	return nil
}
