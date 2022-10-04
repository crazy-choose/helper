package meta

import "github.com/shopspring/decimal"

// 订单信息 && 报单响应
type OrderInfo struct {
	BrokerID             string          `json:"broker_id"`   // 经纪公司代码
	InvestorID           string          `json:"investor_id"` // 投资者代码
	reserve1             string          // 保留的无效字段
	OrderRef             string          `json:"order_ref"`               // 报单引用
	UserID               string          `json:"user_id"`                 // 用户代码
	OrderPriceType       string          `json:"order_price_type"`        // 报单价格条件
	Direction            string          `json:"direction"`               // 买卖方向
	CombOffsetFlag       string          `json:"comb_offset_flag"`        // 组合开平标志
	CombHedgeFlag        string          `json:"comb_hedge_flag"`         // 组合投机套保标志
	LimitPrice           decimal.Decimal `json:"limit_price"`             // 价格
	VolumeTotalOriginal  int             `json:"volume_total_original"`   // 数量
	TimeCondition        string          `json:"time_condition"`          // 有效期类型
	GTDDate              string          `json:"gtd_date"`                // GTD日期
	VolumeCondition      string          `json:"volume_condition"`        // 成交量类型
	MinVolume            int             `json:"min_volume"`              // 最小成交量
	ContingentCondition  string          `json:"contingent_condition"`    // 触发条件
	StopPrice            decimal.Decimal `json:"stop_price"`              // 止损价
	ForceCloseReason     string          `json:"force_close_reason"`      // 强平原因
	IsAutoSuspend        int             `json:"is_auto_suspend"`         // 自动挂起标志
	BusinessUnit         string          `json:"business_unit"`           // 业务单元
	RequestID            int             `json:"request_id"`              // 请求编号
	OrderLocalID         string          `json:"order_local_id"`          // 本地报单编号
	ExchangeID           string          `json:"exchange_id"`             // 交易所代码
	ParticipantID        string          `json:"participant_id"`          // 会员代码
	ClientID             string          `json:"client_id"`               // 客户代码
	reserve2             string          `json:"reserve_2"`               // 保留的无效字段
	TraderID             string          `json:"trader_id"`               // 交易所交易员代码
	InstallID            int             `json:"install_id"`              // 安装编号
	OrderSubmitStatus    string          `json:"order_submit_status"`     // 报单提交状态
	NotifySequence       int             `json:"notify_sequence"`         // 报单提示序号
	TradingDay           string          `json:"trading_day"`             // 交易日
	SettlementID         int             `json:"settlement_id"`           // 结算编号
	OrderSysID           string          `json:"order_sys_id"`            // 报单编号
	OrderSource          string          `json:"order_source"`            // 报单来源
	OrderStatus          string          `json:"order_status"`            // 报单状态
	OrderType            string          `json:"order_type"`              // 报单类型
	VolumeTraded         int             `json:"volume_traded"`           // 今成交数量
	VolumeTotal          int             `json:"volume_total"`            // 剩余数量
	InsertDate           string          `json:"insert_date"`             // 报单日期
	InsertTime           string          `json:"insert_time"`             // 委托时间
	ActiveTime           string          `json:"active_time"`             // 激活时间
	SuspendTime          string          `json:"suspend_time"`            // 挂起时间
	UpdateTime           string          `json:"update_time"`             // 最后修改时间
	CancelTime           string          `json:"cancel_time"`             // 撤销时间
	ActiveTraderID       string          `json:"active_trader_id"`        // 最后修改交易所交易员代码
	ClearingPartID       string          `json:"clearing_part_id"`        // 结算会员编号
	SequenceNo           int             `json:"sequence_no"`             // 序号
	FrontID              int             `json:"front_id"`                // 前置编号
	SessionID            int             `json:"session_id"`              // 会话编号
	UserProductInfo      string          `json:"user_product_info"`       // 用户端产品信息
	StatusMsg            string          `json:"status_msg"`              // 状态信息
	UserForceClose       int             `json:"user_force_close"`        // 用户强评标志
	ActiveUserID         string          `json:"active_user_id"`          // 操作用户代码
	BrokerOrderSeq       int             `json:"broker_order_seq"`        // 经纪公司报单编号
	RelativeOrderSysID   string          `json:"relative_order_sys_id"`   // 相关报单
	ZCETotalTradedVolume int             `json:"zce_total_traded_volume"` // 郑商所成交数量
	IsSwapOrder          int             `json:"is_swap_order"`           // 互换单标志
	BranchID             string          `json:"branch_id"`               // 营业部编号
	InvestUnitID         string          `json:"invest_unit_id"`          // 投资单元代码
	AccountID            string          `json:"account_id"`              // 资金账号
	CurrencyID           string          `json:"currency_id"`             // 币种代码
	reserve3             string          `json:"reserve_3"`               // 保留的无效字段
	MacAddress           string          `json:"mac_address"`             // Mac地址
	InstrumentID         string          `json:"instrument_id"`           // 合约代码
	ExchangeInstID       string          `json:"exchange_inst_id"`        // 合约在交易所的代码
	IPAddress            string          `json:"ip_address"`              // IP地址
	Error                RspInfo         `json:"error"`                   //errorInfo(add)
}

func (*OrderInfo) TableName() string {
	return "order_info"
}

// 报单操作(如：撤单)
type OrderActionInfo struct {
	BrokerID          string          `json:"broker_id"`           // 经纪公司代码
	InvestorID        string          `json:"investor_id"`         // 投资者代码
	OrderActionRef    int             `json:"order_action_ref"`    // 报单操作引用
	OrderRef          string          `json:"order_ref"`           // 报单引用
	RequestID         int             `json:"request_id"`          // 请求编号
	FrontID           int             `json:"front_id"`            // 前置编号
	SessionID         int             `json:"session_id"`          // 会话编号
	ExchangeID        string          `json:"exchange_id"`         // 交易所代码
	OrderSysID        string          `json:"order_sys_id"`        // 报单编号
	ActionFlag        string          `json:"action_flag"`         // 操作标志
	LimitPrice        decimal.Decimal `json:"limit_price"`         // 价格
	VolumeChange      int             `json:"volume_change"`       // 数量变化
	ActionDate        string          `json:"action_date"`         // 操作日期
	ActionTime        string          `json:"action_time"`         // 操作时间
	TraderID          string          `json:"trader_id"`           // 交易所交易员代码
	InstallID         int             `json:"install_id"`          // 安装编号
	OrderLocalID      string          `json:"order_local_id"`      // 本地报单编号
	ActionLocalID     string          `json:"action_local_id"`     // 操作本地编号
	ParticipantID     string          `json:"participant_id"`      // 会员代码
	ClientID          string          `json:"client_id"`           // 客户代码
	BusinessUnit      string          `json:"business_unit"`       // 业务单元
	OrderActionStatus string          `json:"order_action_status"` // 报单操作状态
	UserID            string          `json:"user_id"`             // 用户代码
	StatusMsg         string          `json:"status_msg"`          // 状态信息
	reserve1          string          `json:"reserve_1"`           // 保留的无效字段
	BranchID          string          `json:"branch_id"`           // 营业部编号
	InvestUnitID      string          `json:"invest_unit_id"`      // 投资单元代码
	reserve2          string          `json:"reserve_2"`           // 保留的无效字段
	MacAddress        string          `json:"mac_address"`         // Mac地址
	InstrumentID      string          `json:"instrument_id"`       // 合约代码
	IPAddress         string          `json:"ip_address"`          // IP地址
}

func (*OrderActionInfo) TableName() string {
	return "order_action_info"
}
