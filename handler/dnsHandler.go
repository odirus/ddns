package handler

import (
	"errors"
	"log"
)

type DynamicDns interface {
	// 查询记录，返回值 string -> 记录值, map -> 处理上下文
	queryRecord(domain string, subDomain string, externalIp string) (error, string, *map[string]string)
	// 当没有记录时创建记录
	createRecord(context *map[string]string) error
	// 当存在记录时，更新记录
	updateRecord(context *map[string]string) error
}

type AbstractDynamicDns struct {

}

func (*AbstractDynamicDns) queryRecord(domain string, subDomain string, externalIp string) (error, string, *map[string]string) {
	return errors.New("not implement method"), "", nil
}

func (*AbstractDynamicDns) createRecord(context *map[string]string) error {
	return errors.New("not implement method")
}

func (*AbstractDynamicDns) updateRecord(context *map[string]string) error {
	return errors.New("not implement method")
}

func DoCreateOrUpdateRecord(dns DynamicDns, domain string, subDomain string, externalIp string) {
	err, record, context := dns.queryRecord(domain, subDomain, externalIp)
	if err != nil {
		log.Fatalf("查询记录时出现错误: %s", err.Error())
		return
	}

	// 存在记录时更新记录
	if record != "" {
		// 如果记录未更新则不执行后续请求
		if record == externalIp {
			log.Printf("记录未发生变动, externalIp=%s", externalIp)
			return
		}

		err = dns.updateRecord(context)
		if err != nil {
			log.Fatalf("更新记录时出现错误: %s", err.Error())
			return
		}
		log.Printf("更新记录成功, externalIp=%s", externalIp)
	} else {
		err = dns.createRecord(context)
		if err != nil {
			log.Fatalf("创建记录时出现错误: %s", err.Error())
			return
		}
		log.Printf("创建记录成功, externalIp=%s", externalIp)
	}
}