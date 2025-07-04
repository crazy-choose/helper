package meta

import "github.com/shopspring/decimal"

type KlineType int

const (
	KlineType1Min   KlineType = 1
	KlineType2Min   KlineType = 2
	KlineType3Min   KlineType = 3
	KlineType4Min   KlineType = 4
	KlineType5Min   KlineType = 5
	KlineType6Min   KlineType = 6
	KlineType10Min  KlineType = 10
	KlineType15Min  KlineType = 15
	KlineType30Min  KlineType = 30
	KlineType60Min  KlineType = 60
	KlineType90Min  KlineType = 90
	KlineType120Min KlineType = 120
	KlineType180Min KlineType = 180
	KlineTypeDay    KlineType = 1440
)

type KlineInfo struct {
	Bt            KlineType       `json:"bt"`
	Min           int             `json:"min"`
	IsFinal       bool            `json:"is_final"`
	InstrumentId  string          `json:"instrument_id" gorm:"index:idx_instrument_trading_day,type:btree"` // 创建复合索引，指定类型为 btree
	ExchangeID    string          `json:"exchange_id"`
	TradingDay    string          `json:"trading_day" gorm:"index:idx_instrument_trading_day,type:btree"` // 创建复合索引，指定类型为 btree
	ActionDay     string          `json:"action_day" gorm:"index:idx_action_day,type:btree"`              // 索引
	StartTime     string          `json:"start_time" gorm:"index:idx_start_time,type:btree"`              // 索引
	UpdateTime    string          `json:"update_time" gorm:"index:idx_update_time,type:btree"`            // 索引
	Open          decimal.Decimal `json:"open"`
	High          decimal.Decimal `json:"high"`
	Low           decimal.Decimal `json:"low"`
	Close         decimal.Decimal `json:"close"`
	LastPrice     decimal.Decimal `json:"last_price"`
	SVolume       int             `json:"s_volume"`
	EVolume       int             `json:"e_volume"`
	TurnOver      decimal.Decimal `json:"turn_over"`
	SOpenInterest decimal.Decimal `json:"s_open_interest"`
	EOpenInterest decimal.Decimal `json:"e_open_interest"`
	DisPrice      decimal.Decimal `json:"dis_price"`
	RatioPrice    decimal.Decimal `json:"ratio"`
	DisInterest   decimal.Decimal `json:"dis_interest"`
	Qty           int             `json:"qty"`
}

// KlineInfo K线
type KBInfo struct {
	IsNew         bool            `json:"is_new"`
	Bt            KlineType       `json:"bt"`
	Min           int             `json:"min"`
	IsFinal       bool            `json:"is_final"`
	InstrumentId  string          `json:"instrument_id"`
	ExchangeID    string          `json:"exchange_id"`
	TradingDay    string          `json:"trading_day"` // 交易日
	ActionDay     string          `json:"action_day"`  // 业务日(夜盘品种显示自然日时间)
	StartTime     string          `json:"start_time"`
	UpdateTime    string          `json:"update_time"`
	Open          decimal.Decimal `json:"open"`
	High          decimal.Decimal `json:"high"`
	Low           decimal.Decimal `json:"low"`
	Close         decimal.Decimal `json:"close"`
	LastPrice     decimal.Decimal `json:"last_price"`
	SVolume       int             `json:"s_volume"`        //记录k线初始成交量
	EVolume       int             `json:"e_volume"`        //记录最后tick成交量
	TurnOver      decimal.Decimal `json:"turn_over"`       //成交金额
	SOpenInterest decimal.Decimal `json:"s_open_interest"` //初始持仓量
	EOpenInterest decimal.Decimal `json:"e_open_interest"` //最后持仓量
	DisPrice      decimal.Decimal `json:"dis_price"`       //周期价差
	RatioPrice    decimal.Decimal `json:"ratio"`           //周期涨跌比例
	DisInterest   decimal.Decimal `json:"dis_interest"`    //周期持仓变动
	Qty           int             `json:"qty"`             //周期成交量差值
}
