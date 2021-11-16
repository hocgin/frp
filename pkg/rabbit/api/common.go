package api

import (
	"encoding/json"
	"fmt"
)

type Result struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

/**
 * 处理错误
 */
func handleResult(resultBody string, e error) (data json.RawMessage, err error) {
	if e != nil {
		return nil, e
	}

	result := Result{}
	if err := json.Unmarshal([]byte(resultBody), &result); err != nil {
		err = fmt.Errorf("解析一层数据错误: %v", err)
		return nil, err
	}

	// 请求发生错误
	if !result.Success {
		err = fmt.Errorf("错误信息 [%d], %s", 0, result.Message)
		return
	}
	return result.Data, nil
}
