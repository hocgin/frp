package api

import (
	"encoding/json"
	"fmt"
)

type ProxyChannelInfoVo struct {
	CnvStr string `json:"cnvStr"`
}

func GetChannelInfo(channelId string) (data ProxyChannelInfoVo, err error) {
	path := fmt.Sprintf("/proxy/%s?cnv=frp", channelId)
	dataBody, err := handleResult(UseGet(path, nil))
	if err != nil {
		return
	}

	data = ProxyChannelInfoVo{}
	if err := json.Unmarshal(dataBody, &data); err != nil {
		err = fmt.Errorf("解析二层数据错误: %v", err)
	}
	return
}
