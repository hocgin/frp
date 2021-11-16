package api

import (
	"fmt"
	"github.com/fatedier/frp/pkg/rabbit/config"
	"io"
	"io/ioutil"
	"net/http"
)

type Options struct {
	body io.Reader
}

func UsePost(path string, options *Options) (data string, err error) {
	return UseRequest("POST", path, options)
}

func UseGet(path string, options *Options) (data string, err error) {
	return UseRequest("GET", path, options)
}

func UseRequest(method string, path string, options *Options) (data string, err error) {
	var body io.Reader
	if options != nil {
		body = options.body
	}
	request, err := http.NewRequest(method, rabbitUrl(path), body)
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		err = fmt.Errorf("连接服务器失败: %v", err)
		return
	}
	defer resp.Body.Close()
	resultBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("读取配置数据失败: %v", err)
		return
	}
	return string(resultBody), err
}

func rabbitUrl(path string) string {
	localConfig := config.GetConfig()
	apiserver := localConfig.ApiServer
	return apiserver + path
}
