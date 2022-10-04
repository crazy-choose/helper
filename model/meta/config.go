package meta

//配置
type CTPDeploy struct {
	TradeFront []string
	QuoteFront []string
	BrokerID   string //经纪公司代码
	Password   string
	UserID     string
	AppID      string
	AuthCode   string
	FundPassWD string //资金密码
	DelivDay   int
	Products   []string
}
