package main

import (
	"ddns/handler"
	"ddns/utils"
	"log"
	"os"
	"strings"
	"time"
)

const (
	DnsVendorDNSPod = "DNSPod"
)

func renew() {
	// 记录开始时间
	startMill := time.Now().UnixMilli()

	// 待更新域名
	var domain     = os.Getenv("DDNS_DOMAIN")
	var subDomain  = os.Getenv("DDNS_SUB_DOMAIN")
	var dnsVendor  = os.Getenv("DDNS_DNS_VENDOR")
	var loginToken = os.Getenv("DDNS_DNSPOD_LOGIN_TOKEN")

	// 查询公网IP
	err, externalIp := utils.GetExternalIp()
	if err != nil {
		log.Fatalf("获取公网IP时出现错误: %s", err.Error())
		return
	}

	// 更新记录
	if strings.ToLower(dnsVendor) == strings.ToLower(DnsVendorDNSPod) {
		log.Printf("使用 %s 作为 dns 服务商", DnsVendorDNSPod)
		dnsPodHandler := handler.New(loginToken)
		handler.DoCreateOrUpdateRecord(dnsPodHandler, domain, subDomain, externalIp)
	} else {
		log.Fatalf("不支持的 dns 服务商: %s", dnsVendor)
	}

	// 记录结束时间
	endMill := time.Now().UnixMilli()
	log.Printf("本次任务执行完成, 耗时 %d 毫秒", endMill - startMill)
}

func main() {
	for true {
		renew()
		time.Sleep(3 * time.Minute)
	}
}