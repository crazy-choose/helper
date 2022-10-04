package meta

import "github.com/shopspring/decimal"

// 资金账户
type AccountInfo struct {
	BrokerID                       string          `json:"broker_id"`                           // 经纪公司代码
	AccountID                      string          `json:"account_id"`                          // 投资者帐号
	PreMortgage                    decimal.Decimal `json:"pre_mortgage"`                        // 上次质押金额
	PreCredit                      decimal.Decimal `json:"pre_credit"`                          // 上次信用额度
	PreDeposit                     decimal.Decimal `json:"pre_deposit"`                         // 上次存款额
	PreBalance                     decimal.Decimal `json:"pre_balance"`                         // 上次结算准备金
	PreMargin                      decimal.Decimal `json:"pre_margin"`                          // 上次占用的保证金
	InterestBase                   decimal.Decimal `json:"interest_base"`                       // 利息基数
	Interest                       decimal.Decimal `json:"interest"`                            // 利息收入
	Deposit                        decimal.Decimal `json:"deposit"`                             // 入金金额
	Withdraw                       decimal.Decimal `json:"withdraw"`                            // 出金金额
	FrozenMargin                   decimal.Decimal `json:"frozen_margin"`                       // 冻结的保证金
	FrozenCash                     decimal.Decimal `json:"frozen_cash"`                         // 冻结的资金
	FrozenCommission               decimal.Decimal `json:"frozen_commission"`                   // 冻结的手续费
	CurrMargin                     decimal.Decimal `json:"curr_margin"`                         // 当前保证金总额
	CashIn                         decimal.Decimal `json:"cash_in"`                             // 资金差额
	Commission                     decimal.Decimal `json:"commission"`                          // 手续费
	CloseProfit                    decimal.Decimal `json:"close_profit"`                        // 平仓盈亏
	PositionProfit                 decimal.Decimal `json:"position_profit"`                     // 持仓盈亏
	Balance                        decimal.Decimal `json:"balance"`                             // 期货结算准备金
	Available                      decimal.Decimal `json:"available"`                           // 可用资金
	WithdrawQuota                  decimal.Decimal `json:"withdraw_quota"`                      // 可取资金
	Reserve                        decimal.Decimal `json:"reserve"`                             // 基本准备金
	TradingDay                     string          `json:"trading_day"`                         // 交易日
	SettlementID                   int             `json:"settlement_id"`                       // 结算编号
	Credit                         decimal.Decimal `json:"credit"`                              // 信用额度
	Mortgage                       decimal.Decimal `json:"mortgage"`                            // 质押金额
	ExchangeMargin                 decimal.Decimal `json:"exchange_margin"`                     // 交易所保证金
	DeliveryMargin                 decimal.Decimal `json:"delivery_margin"`                     // 投资者交割保证金
	ExchangeDeliveryMargin         decimal.Decimal `json:"exchange_delivery_margin"`            // 交易所交割保证金
	ReserveBalance                 decimal.Decimal `json:"reserve_balance"`                     // 保底期货结算准备金
	CurrencyID                     string          `json:"currency_id"`                         // 币种代码
	PreFundMortgageIn              decimal.Decimal `json:"pre_fund_mortgage_in"`                // 上次货币质入金额
	PreFundMortgageOut             decimal.Decimal `json:"pre_fund_mortgage_out"`               // 上次货币质出金额
	FundMortgageIn                 decimal.Decimal `json:"fund_mortgage_in"`                    // 货币质入金额
	FundMortgageOut                decimal.Decimal `json:"fund_mortgage_out"`                   // 货币质出金额
	FundMortgageAvailable          decimal.Decimal `json:"fund_mortgage_available"`             // 货币质押余额
	MortgageableFund               decimal.Decimal `json:"mortgageable_fund"`                   // 可质押货币金额
	SpecProductMargin              decimal.Decimal `json:"spec_product_margin"`                 // 特殊产品占用保证金
	SpecProductFrozenMargin        decimal.Decimal `json:"spec_product_frozen_margin"`          // 特殊产品冻结保证金
	SpecProductCommission          decimal.Decimal `json:"spec_product_commission"`             // 特殊产品手续费
	SpecProductFrozenCommission    decimal.Decimal `json:"spec_product_frozen_commission"`      // 特殊产品冻结手续费
	SpecProductPositionProfit      decimal.Decimal `json:"spec_product_position_profit"`        // 特殊产品持仓盈亏
	SpecProductCloseProfit         decimal.Decimal `json:"spec_product_close_profit"`           // 特殊产品平仓盈亏
	SpecProductPositionProfitByAlg decimal.Decimal `json:"spec_product_position_profit_by_alg"` // 根据持仓盈亏算法计算的特殊产品持仓盈亏
	SpecProductExchangeMargin      decimal.Decimal `json:"spec_product_exchange_margin"`        // 特殊产品交易所保证金
	BizType                        string          `json:"biz_type"`                            // 业务类型
	FrozenSwap                     decimal.Decimal `json:"frozen_swap"`                         // 延时换汇冻结金额
	RemainSwap                     decimal.Decimal `json:"remain_swap"`                         // 剩余换汇额度
}

func (*AccountInfo) TableName() string {
	return "account_info"
}

//持仓
type Position struct {
	Reserve1           string          `json:"reserve1"`              //保留的无效字段
	BrokerID           string          `json:"reserve1"`              //经纪公司代码
	InvestorID         string          `json:"reserve1"`              //投资者代码
	PosiDirection      string          `json:"posi_direction"`        //持仓多空方向
	HedgeFlag          string          `json:"hedge_flag"`            // 投机套保标志
	PositionDate       string          `json:"position_date"`         // 持仓日期
	Position           int             `json:"position"`              // 当前总持仓
	YdPosition         int             `json:"yd_position"`           // 上日持仓
	TodayPosition      int             `json:"today_position"`        // 表示今新开仓
	OpenVolume         int             `json:"open_volume"`           // 开仓量
	CloseVolume        int             `json:"close_volume"`          // 平仓量
	OpenAmount         decimal.Decimal `json:"open_amount"`           // 开仓金额
	CloseAmount        decimal.Decimal `json:"close_amount"`          // 平仓金额
	PositionCost       decimal.Decimal `json:"position_cost"`         // 持仓成本=上日持仓 * 昨结算价 * 合约乘数 + SUM（今日持仓 * 开仓价 * 合约乘数）
	CloseProfit        decimal.Decimal `json:"close_profit"`          // 平仓盈亏
	PositionProfit     decimal.Decimal `json:"position_profit"`       // 持仓盈亏
	OpenCost           decimal.Decimal `json:"open_cost"`             // 开仓成本
	LongFrozen         int             `json:"long_frozen"`           // 多头冻结
	ShortFrozen        int             `json:"short_frozen"`          // 空头冻结
	LongFrozenAmount   decimal.Decimal `json:"long_frozen_amount"`    // 开仓冻结金额
	ShortFrozenAmount  decimal.Decimal `json:"short_frozen_amount"`   // 开仓冻结金额
	PreMargin          decimal.Decimal `json:"pre_margin"`            // 上次占用的保证金
	UseMargin          decimal.Decimal `json:"use_margin"`            // 占用的保证金
	FrozenMargin       decimal.Decimal `json:"frozen_margin"`         // 冻结的保证金
	FrozenCash         decimal.Decimal `json:"frozen_cash"`           // 冻结的资金
	FrozenCommission   decimal.Decimal `json:"frozen_commission"`     // 冻结的手续费
	CashIn             decimal.Decimal `json:"cash_in"`               // 资金差额
	Commission         decimal.Decimal `json:"commission"`            // 手续费
	PreSettlementPrice decimal.Decimal `json:"pre_settlement_price"`  // 上次结算价
	SettlementPrice    decimal.Decimal `json:"settlement_price"`      // 本次结算价
	SettlementID       int             `json:"settlement_id"`         //结算编号
	ExchangeMargin     decimal.Decimal `json:"exchange_margin"`       // 交易所保证金
	CombPosition       int             `json:"comb_position"`         // 组合成交形成的持仓
	CombLongFrozen     int             `json:"comb_long_frozen"`      // 组合多头冻结
	CombShortFrozen    int             `json:"comb_short_frozen"`     // 组合空头冻结
	CloseProfitByDate  decimal.Decimal `json:"close_profit_by_date"`  // 逐日盯市平仓盈亏
	CloseProfitByTrade decimal.Decimal `json:"close_profit_by_trade"` // 逐笔对冲平仓盈亏
	StrikeFrozen       int             `json:"strike_frozen"`         // 执行冻结
	StrikeFrozenAmount decimal.Decimal `json:"strike_frozen_amount"`  // 执行冻结金额
	AbandonFrozen      int             `json:"abandon_frozen"`        // 放弃执行冻结
	ExchangeID         string          `json:"exchange_id"`           // 交易所代码
	YdStrikeFrozen     int             `json:"yd_strike_frozen"`      // 执行冻结的昨仓
	PositionCostOffset decimal.Decimal `json:"position_cost_offset"`  // 大商所持仓成本差值，只有大商所使用
	InstrumentID       string          `json:"instrument_id"`         // 合约代码
	InstrumentName     string          `json:"instrument_name"`       // 合约名
}

func (*Position) TableName() string {
	return "position"
}
