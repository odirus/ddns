package handler

import (
	"ddns/dnspod"
	"errors"
	"fmt"
)



type DNSPodHandler struct {
	dnsPodConf *dnspod.Conf
	AbstractDynamicDns
}

func New(loginToken string) *DNSPodHandler {
	handler := &DNSPodHandler{
		dnsPodConf: &dnspod.Conf{LoginToken: loginToken},
	}
	return handler
}

func (handler *DNSPodHandler) queryRecord(domain string, subDomain string, externalIp string) (error, string, *map[string]string) {
	criteria := &dnspod.RecordCriteria {
		Domain: domain,
		SubDomain: subDomain,
		RecordType: "A",
	}
	err, resp := dnspod.List(handler.dnsPodConf, criteria)
	if err != nil {
		return errors.New(fmt.Sprintf("通过DNSPod API查询记录时出现错误: %s", err.Error())), "", nil
	}
	if len(resp.Records) == 0 {
		return nil, "", &map[string]string{
			"domain": domain,
			"subDomain": subDomain,
			"externalIp": externalIp,
		}
	} else if len(resp.Records) > 1 {
		return errors.New(fmt.Sprintf("通过DNSPod API查询到超过一条记录值: %s", resp.Records)), "", nil
	} else {
		return nil, resp.Records[0].Value, &map[string]string{
			"domain": domain,
			"subDomain": subDomain,
			"recordId": resp.Records[0].Id,
			"externalIp": externalIp,
		}
	}
}

func (handler *DNSPodHandler) createRecord(context *map[string]string) error {
	criteria := &dnspod.RecordCriteria {
		Domain: (*context)["domain"],
		SubDomain: (*context)["subDomain"],
		RecordType: "A",
		RecordLine: "默认",
		Value: (*context)["externalIp"],
	}
	err := dnspod.Create(handler.dnsPodConf, criteria)
	if err != nil {
		return errors.New(fmt.Sprintf("通过DNSPod API创建记录时出现错误: %s", err.Error()))
	} else {
		return nil
	}
}

func (handler *DNSPodHandler) updateRecord(context *map[string]string) error {
	criteria := &dnspod.RecordCriteria {
		Domain: (*context)["domain"],
		SubDomain: (*context)["subDomain"],
		RecordType: "A",
		RecordId: (*context)["recordId"],
		RecordLine: "默认",
		Value: (*context)["externalIp"],
	}
	err := dnspod.Modify(handler.dnsPodConf, criteria)
	if err != nil {
		return errors.New(fmt.Sprintf("通过DNSPod API修改记录时出现错误: %s", err.Error()))
	} else {
		return nil
	}
}
