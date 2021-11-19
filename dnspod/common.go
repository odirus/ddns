package dnspod

const (
	ApiRecordList = "https://dnsapi.cn/Record.List"
	ApiRecordModify = "https://dnsapi.cn/Record.Modify"
	ApiRecordCreate = "https://dnsapi.cn/Record.Create"
)

const (
	ApiRespCodeSuccess = "1"
)

type Conf struct {
	LoginToken string
}

type ApiRespStatus struct {
	// code, 1 success
	Code string `json:"code"`
	// message
	Message string `json:"message"`
}

type CommonResp struct {
	Status ApiRespStatus `json:"status"`
}