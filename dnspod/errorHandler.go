package dnspod

import (
	"bytes"
	"fmt"
	"github.com/go-resty/resty/v2"
	"reflect"
)

func FormatHttpRequestErr(resp *resty.Response, err error) string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("请求HTTP接口时出现错误:"))
	buffer.WriteString(fmt.Sprintf("  Error: %s", err.Error()))
	buffer.WriteString(fmt.Sprintf("  Status Code: %d", resp.StatusCode()))
	buffer.WriteString(fmt.Sprintf("  Time: %s", resp.Time()))
	buffer.WriteString(fmt.Sprintf("  Received At: %s", resp.ReceivedAt()))
	buffer.WriteString(fmt.Sprintf("  Body: %s", resp.Body()))
	return buffer.String()
}

func MarshalErr(marshalType interface{}, err error) string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("序列化JSON是出现错误:"))
	buffer.WriteString(fmt.Sprintf("  Error: %s", err.Error()))
	buffer.WriteString(fmt.Sprintf("  Type: %s", reflect.ValueOf(marshalType).String()))
	return buffer.String()
}

func UnmarshalErr(unmarshalType interface{}, response string, err error) string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("反序列化JSON是出现错误:"))
	buffer.WriteString(fmt.Sprintf("  Error: %s", err.Error()))
	buffer.WriteString(fmt.Sprintf("  Response: %s", response))
	buffer.WriteString(fmt.Sprintf("  Type: %s", reflect.ValueOf(unmarshalType).String()))
	return buffer.String()
}

func BizError(status ApiRespStatus) string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Dnspod 业务接口调用失败:"))
	buffer.WriteString(fmt.Sprintf("  Status.Code: %s", status.Code))
	buffer.WriteString(fmt.Sprintf("  Status.Message: %s", status.Message))
	return buffer.String()
}
