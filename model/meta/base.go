package meta

// 响应信息
type RspInfo struct {
	ErrorID  int    `json:"error_id"`  // 错误代码
	ErrorMsg string `json:"error_msg"` // 错误信息
}

// 用户登录应答
type RspUserLogin struct {
	TradingDay  string `json:"trading_day"`   // 交易日
	LoginTime   string `json:"login_time"`    // 登录成功时间
	BrokerID    string `json:"broker_id"`     // 经纪公司代码
	UserID      string `json:"user_id"`       // 用户代码
	SystemName  string `json:"system_name"`   // 交易系统名称
	FrontID     int    `json:"front_id"`      // 前置编号
	SessionID   int    `json:"session_id"`    // 会话编号
	MaxOrderRef string `json:"max_order_ref"` // 最大报单引用
	SHFETime    string `json:"shfe_time"`     // 上期所时间
	DCETime     string `json:"dce_time"`      // 大商所时间
	CZCETime    string `json:"czce_time"`     // 郑商所时间
	FFEXTime    string `json:"ffex_time"`     // 中金所时间
	INETime     string `json:"ine_time"`      // 能源中心时间
	SysVersion  string `json:"sys_version"`   //后台版本信息
}

type CallBackInfo struct {
	Field     interface{} //msg Field
	Rsp       *RspInfo
	RequestId int
	IsLast    bool
}
