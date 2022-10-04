package meta

import "github.com/shopspring/decimal"

//成交
type TradeInfo struct {
	BrokerID       string          `json:"broker_id"`        // 经纪公司代码
	InvestorID     string          `json:"investor_id"`      // 投资者代码
	reserve1       string          `json:"reserve_1"`        // 保留的无效字段
	OrderRef       string          `json:"order_ref"`        // 报单引用
	UserID         string          `json:"user_id"`          // 用户代码
	ExchangeID     string          `json:"exchange_id"`      // 交易所代码
	TradeID        string          `json:"trade_id"`         // 成交编号
	Direction      string          `json:"direction"`        // 买卖方向
	OrderSysID     string          `json:"order_sys_id"`     // 报单编号
	ParticipantID  string          `json:"participant_id"`   // 会员代码
	ClientID       string          `json:"client_id"`        // 客户代码
	TradingRole    string          `json:"trading_role"`     // 交易角色
	reserve2       string          `json:"reserve_2"`        // 保留的无效字段
	OffsetFlag     string          `json:"offset_flag"`      // 开平标志
	HedgeFlag      string          `json:"hedge_flag"`       // 投机套保标志
	Price          decimal.Decimal `json:"price"`            // 价格
	Volume         int             `json:"volume"`           // 数量
	TradeDate      string          `json:"trade_date"`       // 成交时期
	TradeTime      string          `json:"trade_time"`       // 成交时间
	TradeType      string          `json:"trade_type"`       // 成交类型
	PriceSource    string          `json:"price_source"`     // 成交价来源
	TraderID       string          `json:"trader_id"`        // 交易所交易员代码
	OrderLocalID   string          `json:"order_local_id"`   // 本地报单编号
	ClearingPartID string          `json:"clearing_part_id"` // 结算会员编号
	BusinessUnit   string          `json:"business_unit"`    // 业务单元
	SequenceNo     int             `json:"sequence_no"`      // 序号
	TradingDay     string          `json:"trading_day"`      // 交易日
	SettlementID   int             `json:"settlement_id"`    // 结算编号
	BrokerOrderSeq int             `json:"broker_order_seq"` // 经纪公司报单编号
	TradeSource    string          `json:"trade_source"`     // 成交来源
	InvestUnitID   string          `json:"invest_unit_id"`   // 投资单元代码
	InstrumentID   string          `json:"instrument_id"`    // 合约代码
	ExchangeInstID string          `json:"exchange_inst_id"` // 合约在交易所的代码
}

func (*TradeInfo) TableName() string {
	return "trade_info"
}

// 合约信息
type InstrumentInfo struct {
	reserve1               string          `json:"reserve_1"`                 // 保留的无效字段
	ExchangeID             string          `json:"exchange_id"`               // 交易所代码
	InstrumentName         string          `json:"instrument_name"`           // 合约名称
	reserve2               string          `json:"reserve_2"`                 // 保留的无效字段
	reserve3               string          `json:"reserve_3"`                 // 保留的无效字段
	ProductClass           string          `json:"product_class"`             // 产品类型
	DeliveryYear           int             `json:"delivery_year"`             // 交割年份
	DeliveryMonth          int             `json:"delivery_month"`            // 交割月
	MaxMarketOrderVolume   int             `json:"max_market_order_volume"`   // 市价单最大下单量
	MinMarketOrderVolume   int             `json:"min_market_order_volume"`   // 市价单最小下单量
	MaxLimitOrderVolume    int             `json:"max_limit_order_volume"`    // 限价单最大下单量
	MinLimitOrderVolume    int             `json:"min_limit_order_volume"`    // 限价单最小下单量
	VolumeMultiple         int             `json:"volume_multiple"`           // 合约数量乘数
	PriceTick              decimal.Decimal `json:"price_tick"`                // 最小变动价位
	CreateDate             string          `json:"create_date"`               // 创建日
	OpenDate               string          `json:"open_date"`                 // 上市日
	ExpireDate             string          `json:"expire_date"`               // 到期日
	StartDelivDate         string          `json:"start_deliv_date"`          // 开始交割日
	EndDelivDate           string          `json:"end_deliv_date"`            // 结束交割日
	InstLifePhase          string          `json:"inst_life_phase"`           // 合约生命周期状态
	IsTrading              int             `json:"is_trading"`                // 当前是否交易
	PositionType           string          `json:"position_type"`             // 持仓类型
	PositionDateType       string          `json:"position_date_type"`        // 持仓日期类型
	LongMarginRatio        decimal.Decimal `json:"long_margin_ratio"`         // 多头保证金率
	ShortMarginRatio       decimal.Decimal `json:"short_margin_ratio"`        // 空头保证金率
	MaxMarginSideAlgorithm string          `json:"max_margin_side_algorithm"` // 是否使用大额单边保证金算法
	reserve4               string          `json:"reserve_4"`                 // 保留的无效字段
	StrikePrice            decimal.Decimal `json:"strike_price"`              // 执行价
	OptionsType            string          `json:"options_type"`              // 期权类型
	UnderlyingMultiple     decimal.Decimal `json:"underlying_multiple"`       // 合约基础商品乘数
	CombinationType        string          `json:"combination_type"`          // 组合类型
	InstrumentID           string          `json:"instrument_id"`             // 合约代码
	ExchangeInstID         string          `json:"exchange_inst_id"`          // 合约在交易所的代码
	ProductID              string          `json:"product_id"`                // 产品代码
	UnderlyingInstrID      string          `json:"underlying_instr_id"`       // 基础商品代码
}

func (*InstrumentInfo) TableName() string {
	return "instrument_info"
}

// InstrumentStatus 合约状态
type InstrumentStatus struct {
	ExchangeID        string `json:"exchange_id"`         // 交易所代码	[9]byte
	reserve1          string `json:"-"`                   // 保留的无效字段	[31]byte
	SettlementGroupID string `json:"settlement_group_id"` // 结算组代码 [9]byte
	reserve2          string `json:"-"`                   // 保留的无效字段 [31]byte
	InstrumentStatus  string `json:"instrument_status"`   // 合约交易状态 byte
	TradingSegmentSN  int    `json:"trading_segment_sn"`  // 交易阶段编号 int32
	EnterTime         string `json:"enter_time"`          // 进入本状态时间 [9]byte
	EnterReason       string `json:"enter_reason"`        // 进入本状态原因 byte
	ExchangeInstID    string `json:"exchange_inst_id"`    // 合约在交易所的代码 [81]byte
	InstrumentID      string `json:"instrument_id"`       // 合约代码 [81]byte
	StatusZN          string `json:"status_zn"`           //当前状态中文

	PreTradingSegmentSN int    `json:"pre_trading_segment_sn"` // 前交易阶段编号(自加)
	PreStatus           string `json:"pre_status"`             //前状态(自加)
	PreEnterTime        string `json:"pre_enter_time"`         //前状态进入时间(自加)
	PreEnterReason      string `json:"pre_enter_reason"`       //前状态进入原因(自加)
}

func (*InstrumentStatus) TableName() string {
	return "instrument_status"
}
