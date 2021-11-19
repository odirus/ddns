# README
通过DNS API实现 DDNS 功能，目前支持的DNS服务商有 DNSPod 等

## 打包命令
```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o linux-amd64-ddns main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o windows-amd64-ddns.exe main.go
```

## 启动参数
参数通过环境变量传入，目前支持的环境变量有：
* DDNS_DOMAIN                  顶级域名，例如 odirus.me  
* DDNS_SUB_DOMAIN              二级域名，例如 home  
* DDNS_DNS_VENDOR              DNS服务商，目前只支持 DNSPod  
* DDNS_DNSPOD_LOGIN_TOKEN      DNSPod的LoginToken，获取方式参考[DNSPod密钥管理](https://docs.dnspod.cn/account/5f2d466de8320f1a740d9ff3/)
