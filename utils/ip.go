package utils

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
)

const (
	apiGetExternalIp = "https://myexternalip.com/raw"
)

func GetExternalIp() (error, string) {
	client := resty.New()
	resp, err := client.R().Get(apiGetExternalIp)
	if err != nil {
		return errors.New(fmt.Sprintf("通过 myexternalip 查询 externalIp 时出现错误: %s", err.Error())), ""
	}

	if resp.StatusCode() == 200 {
		return nil, resp.String()
	} else {
		return errors.New(fmt.Sprintf("通过 myexternalip 查询 externalIp 时出现错误, HTTP statusCode: %d", resp.StatusCode())), ""
	}
}