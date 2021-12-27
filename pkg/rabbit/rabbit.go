package rabbit

import (
	"fmt"
	"github.com/fatedier/frp/pkg/rabbit/api"
	"io/ioutil"
)

/**
 * 从服务端拉取配置信息
 */
func DownloadClientConfig4Server(channelId string) (
	cfgFilePath string,
	err error,
) {
	// 1. 获取配置信息
	data, err := api.GetChannelInfo(channelId)
	if err != nil {
		err = fmt.Errorf("获取隧道[id=%s]信息错误, error: %v", channelId, err)
		return
	}

	// 2. 保存配置文件
	cfgFilePath, err = saveRabbit2Cfg(data.CnvStr)
	if err != nil {
		err = fmt.Errorf("保存隧道[id=%s]信息错误, error: %v", channelId, err)
		return
	}
	return
}

func saveRabbit2Cfg(body string) (
	cfgPath string,
	err error,
) {

	// 1. 解析 json

	// 2. md5 校验

	// 3. 转换为 ini

	// 4. 存储到本地
	file, err := ioutil.TempFile("", "frpc.ini")
	if err != nil {
		return
	}
	_, err = file.WriteString(body)
	if err != nil {
		return
	}
	cfgPath = file.Name()
	// fmt.Println("--> 临时配置存储位置: " + cfgPath)
	return
}
