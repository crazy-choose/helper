package meta

import "github.com/shopspring/decimal"

// 成交
type TradeInfo struct {
	BrokerID       string          `json:"broker_id"`        // 经纪公司代码
	InvestorID     string          `json:"investor_id"`      // 投资者代码
	reserve1       string          `json:"reserve_1"`        // 保留的无效字段
	OrderRef       string          `json:"order_ref"`        // 报单引用
	UserID         string          `json:"user_id"`          // 用户代码
	ExchangeID     string          `json:"exchange_id"`      // 交易所代码
	TradeID        string          `json:"trade_id"`         // 成交编号
	Direction      int32           `json:"direction"`        // 买卖方向
	OrderSysID     string          `json:"order_sys_id"`     // 报单编号
	ParticipantID  string          `json:"participant_id"`   // 会员代码
	ClientID       string          `json:"client_id"`        // 客户代码
	TradingRole    string          `json:"trading_role"`     // 交易角色
	reserve2       string          `json:"reserve_2"`        // 保留的无效字段
	OffsetFlag     int32           `json:"offset_flag"`      // 开平标志
	HedgeFlag      string          `json:"hedge_flag"`       // 投机套保标志
	Price          decimal.Decimal `json:"price"`            // 价格
	Volume         int             `json:"volume"`           // 数量
	TradeDate      string          `json:"trade_date"`       // 成交时期
	TradeTime      string          `json:"trade_time"`       // 成交时间
	TradeType      int32           `json:"trade_type"`       // 成交类型
	PriceSource    string          `json:"price_source"`     // 成交价来源
	TraderID       string          `json:"trader_id"`        // 交易所交易员代码
	OrderLocalID   string          `json:"order_local_id"`   // 本地报单编号
	ClearingPartID string          `json:"clearing_part_id"` // 结算会员编号
	BusinessUnit   string          `json:"business_unit"`    // 业务单元
	SequenceNo     int             `json:"sequence_no"`      // 序号
	TradingDay     string          `json:"trading_day"`      // 交易日
	SettlementID   int             `json:"settlement_id"`    // 结算编号
	BrokerOrderSeq int             `json:"broker_order_seq"` // 经纪公司报单编号
	TradeSource    int32           `json:"trade_source"`     // 成交来源
	InvestUnitID   string          `json:"invest_unit_id"`   // 投资单元代码
	InstrumentID   string          `json:"instrument_id"`    // 合约代码
	ExchangeInstID string          `json:"exchange_inst_id"` // 合约在交易所的代码
}
