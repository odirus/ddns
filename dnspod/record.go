package dnspod

import (
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
)

type RecordCriteria struct {
	// 域名
	Domain string `json:"domain"`
	// 子域名
	SubDomain string `json:"sub_domain"`
	// 记录类型
	RecordType string `json:"record_type"`
	// 记录ID编号
	RecordId string `json:"record_id"`
	// 记录线路，推荐 "默认"
	RecordLine string `json:"record_line"`
	// 记录值
	Value string `json:"value"`
}

type Record struct {
	// 记录ID编号
	Id string `json:"id"`
	// 记录值
	Value string `json:"value"`
}

type ListRecordResp struct {
	Records []Record `json:"records"`
}

func convertStruct2Map(v interface{}) (error, *map[string]string) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return errors.New(MarshalErr(v, err)), nil
	}

	data := &map[string]string{}
	err = json.Unmarshal(bytes, data)
	if err != nil {
		return errors.New(UnmarshalErr(data, string(bytes), err)), nil
	}

	return nil, data
}

func doRequest(conf *Conf, endpoint string, criteria interface{}, resp interface{}) error {
	requestBody := &map[string]string{}
	// 合并基础参数
	apiBaseParam := &map[string]string{
		"login_token": conf.LoginToken,
		"format": "json",
		"error_on_empty": "no",
	}
	for k, v := range *apiBaseParam {
		(*requestBody)[k] = v
	}
	// 合并业务参数
	convertErr, requestParamMap := convertStruct2Map(criteria)
	if convertErr != nil {
		return convertErr
	}
	for k, v := range *requestParamMap {
		(*requestBody)[k] = v
	}

	// 执行请求
	client := resty.New()
	apiResp, err := client.R().
		SetFormData(*requestBody).Post(endpoint)
	if err != nil {
		return errors.New(FormatHttpRequestErr(apiResp, err))
	}

	// 解析公共结果
	commonResp := &CommonResp{}
	err = json.Unmarshal(apiResp.Body(), commonResp)
	if err != nil {
		return errors.New(UnmarshalErr(resp, apiResp.String(), err))
	}
	// 判断业务状态码
	if commonResp.Status.Code != ApiRespCodeSuccess {
		return errors.New(BizError(commonResp.Status))
	}

	// 解析业务结果
	if resp != nil {
		err = json.Unmarshal(apiResp.Body(), resp)
		if err != nil {
			return errors.New(UnmarshalErr(resp, apiResp.String(), err))
		}
	}

	return nil
}

func List(conf *Conf, criteria *RecordCriteria) (err error, resp *ListRecordResp) {
	listRecordResp := &ListRecordResp{}
	err = doRequest(conf, ApiRecordList, criteria, listRecordResp)
	if err != nil {
		return err, nil
	}

	return nil, listRecordResp
}

func Modify(conf *Conf, criteria *RecordCriteria) error {
	err := doRequest(conf, ApiRecordModify, criteria, nil)
	if err != nil {
		return err
	}

	return nil
}

func Create(conf *Conf, criteria *RecordCriteria) error {
	err := doRequest(conf, ApiRecordCreate, criteria, nil)
	if err != nil {
		return err
	}

	return nil
}