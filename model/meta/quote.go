package meta

import "github.com/shopspring/decimal"

type TickCache struct {
	TradingDay         string          `json:"trading_day"`          // 交易日
	InstrumentID       string          `json:"instrument_id"`        // 合约代码
	ExchangeID         string          `json:"exchange_id"`          // 交易所代码
	ExchangeInstID     string          `json:"exchange_inst_id"`     // 合约在交易所的代码
	LastPrice          decimal.Decimal `json:"last_price"`           // 最新价
	PreSettlementPrice decimal.Decimal `json:"pre_settlement_price"` // 上次结算价
	PreClosePrice      decimal.Decimal `json:"pre_close_price"`      // 昨收盘
	PreOpenInterest    decimal.Decimal `json:"pre_open_interest"`    // 昨持仓量
	OpenPrice          decimal.Decimal `json:"open_price"`           // 今开盘
	HighestPrice       decimal.Decimal `json:"highest_price"`        // 最高价
	LowestPrice        decimal.Decimal `json:"lowest_price"`         // 最低价
	Volume             int             `json:"volume"`               // 数量
	Turnover           decimal.Decimal `json:"turnover"`             // 成交金额
	OpenInterest       decimal.Decimal `json:"open_interest"`        // 持仓量
	ClosePrice         decimal.Decimal `json:"close_price"`          // 今收盘
	SettlementPrice    decimal.Decimal `json:"settlement_price"`     // 本次结算价
	UpperLimitPrice    decimal.Decimal `json:"upper_limit_price"`    // 涨停板价
	LowerLimitPrice    decimal.Decimal `json:"lower_limit_price"`    // 跌停板价
	PreDelta           decimal.Decimal `json:"pre_delta"`            // 昨虚实度
	CurrDelta          decimal.Decimal `json:"curr_delta"`           // 今虚实度
	UpdateTime         string          `json:"update_time"`          // 最后修改时间
	UpdateMillisec     int             `json:"update_millisec"`      // 最后修改毫秒
	AveragePrice       decimal.Decimal `json:"average_price"`        // 当日均价
	ActionDay          string          `json:"action_day"`           // 业务日期
}

// 行情响应
type TickInfo struct {
	TickCache
	BidPrice1         decimal.Decimal `json:"bid_price_1"`         // 申买价一
	BidVolume1        int             `json:"bid_volume_1"`        // 申买量一
	AskPrice1         decimal.Decimal `json:"ask_price_1"`         // 申卖价一
	AskVolume1        int             `json:"ask_volume_1"`        // 申卖量一
	BidPrice2         decimal.Decimal `json:"bid_price_2"`         // 申买价二
	BidVolume2        int             `json:"bid_volume_2"`        // 申买量二
	AskPrice2         decimal.Decimal `json:"ask_price_2"`         // 申卖价二
	AskVolume2        int             `json:"ask_volume_2"`        // 申卖量二
	BidPrice3         decimal.Decimal `json:"bid_price_3"`         // 申买价三
	BidVolume3        int             `json:"bid_volume_3"`        // 申买量三
	AskPrice3         decimal.Decimal `json:"ask_price_3"`         // 申卖价三
	AskVolume3        int             `json:"ask_volume_3"`        // 申卖量三
	BidPrice4         decimal.Decimal `json:"bid_price_4"`         // 申买价四
	BidVolume4        int             `json:"bid_volume_4"`        // 申买量四
	AskPrice4         decimal.Decimal `json:"ask_price_4"`         // 申卖价四
	AskVolume4        int             `json:"ask_volume_4"`        // 申卖量四
	BidPrice5         decimal.Decimal `json:"bid_price_5"`         // 申买价五
	BidVolume5        int             `json:"bid_volume_5"`        // 申买量五
	AskPrice5         decimal.Decimal `json:"ask_price_5"`         // 申卖价五
	AskVolume5        int             `json:"ask_volume_5"`        // 申卖量五
	BandingUpperPrice decimal.Decimal `json:"banding_upper_price"` // 上待价
	BandingLowerPrice decimal.Decimal `json:"banding_lower_price"` // 下待价
}

type SpecificInstrument struct {
	InstrumentID string `json:"instrument_id"` // 合约代码
}
