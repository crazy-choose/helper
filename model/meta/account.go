package meta

import "github.com/shopspring/decimal"

// 资金账户
type AccountInfo struct {
	BrokerID                       string          `json:"broker_id" gorm:"column:broker_id;comment:经纪公司代码;"`                                                                  // 经纪公司代码
	AccountID                      string          `json:"account_id" gorm:"column:account_id;comment:投资者帐号;index:idx_account_id,type:btree"`                                  // 投资者帐号
	PreMortgage                    decimal.Decimal `json:"pre_mortgage" gorm:"column:pre_mortgage;comment:上次质押金额;"`                                                            // 上次质押金额
	PreCredit                      decimal.Decimal `json:"pre_credit" gorm:"column:pre_credit;comment:上次信用额度;"`                                                                // 上次信用额度
	PreDeposit                     decimal.Decimal `json:"pre_deposit" gorm:"column:pre_deposit;comment:上次存款额;"`                                                               // 上次存款额
	PreBalance                     decimal.Decimal `json:"pre_balance" gorm:"column:pre_balance;comment:上次结算准备金;"`                                                             // 上次结算准备金
	PreMargin                      decimal.Decimal `json:"pre_margin" gorm:"column:pre_margin;comment:上次占用的保证金;"`                                                              // 上次占用的保证金
	InterestBase                   decimal.Decimal `json:"interest_base" gorm:"column:interest_base;comment:利息基数;"`                                                            // 利息基数
	Interest                       decimal.Decimal `json:"interest" gorm:"column:interest;comment:利息收入;"`                                                                      // 利息收入
	Deposit                        decimal.Decimal `json:"deposit" gorm:"column:deposit;comment:入金金额;"`                                                                        // 入金金额
	Withdraw                       decimal.Decimal `json:"withdraw" gorm:"column:withdraw;comment:出金金额;"`                                                                      // 出金金额
	FrozenMargin                   decimal.Decimal `json:"frozen_margin" gorm:"column:frozen_margin;comment:冻结的保证金;"`                                                          // 冻结的保证金
	FrozenCash                     decimal.Decimal `json:"frozen_cash" gorm:"column:frozen_cash;comment:冻结的资金;"`                                                               // 冻结的资金
	FrozenCommission               decimal.Decimal `json:"frozen_commission" gorm:"column:frozen_commission;comment:冻结的手续费;"`                                                  // 冻结的手续费
	CurrMargin                     decimal.Decimal `json:"curr_margin" gorm:"column:curr_margin;comment:当前保证金总额;"`                                                             // 当前保证金总额
	CashIn                         decimal.Decimal `json:"cash_in" gorm:"column:cash_in;comment:资金差额;"`                                                                        // 资金差额
	Commission                     decimal.Decimal `json:"commission" gorm:"column:commission;comment:手续费;"`                                                                   // 手续费
	CloseProfit                    decimal.Decimal `json:"close_profit" gorm:"column:close_profit;comment:平仓盈亏;index:idx_close_profit,type:btree"`                             // 平仓盈亏
	PositionProfit                 decimal.Decimal `json:"position_profit" gorm:"column:position_profit;comment:持仓盈亏;"`                                                        // 持仓盈亏
	Balance                        decimal.Decimal `json:"balance" gorm:"column:balance;comment:期货结算准备金;"`                                                                     // 期货结算准备金
	Available                      decimal.Decimal `json:"available" gorm:"column:available;comment:可用资金;index:idx_available,type:btree"`                                      // 可用资金
	WithdrawQuota                  decimal.Decimal `json:"withdraw_quota" gorm:"column:withdraw_quota;comment:期货结算准备金;"`                                                       // 可取资金
	Reserve                        decimal.Decimal `json:"reserve" gorm:"column:reserve;comment:期货结算准备金;"`                                                                     // 基本准备金
	TradingDay                     string          `json:"trading_day" gorm:"column:trading_day;comment:交易日;index:idx_trading_day,type:btree"`                                 // 交易日
	SettlementID                   int             `json:"settlement_id" gorm:"column:settlement_id;comment:结算编号;index:idx_settlement_id,type:btree"`                          // 结算编号
	Credit                         decimal.Decimal `json:"credit" gorm:"column:credit;comment:信用额度;"`                                                                          // 信用额度
	Mortgage                       decimal.Decimal `json:"mortgage" gorm:"column:mortgage;comment:质押金额;"`                                                                      // 质押金额
	ExchangeMargin                 decimal.Decimal `json:"exchange_margin" gorm:"column:exchange_margin;comment:交易所保证金;"`                                                      // 交易所保证金
	DeliveryMargin                 decimal.Decimal `json:"delivery_margin" gorm:"column:delivery_margin;comment:投资者交割保证金;"`                                                    // 投资者交割保证金
	ExchangeDeliveryMargin         decimal.Decimal `json:"exchange_delivery_margin" gorm:"column:exchange_delivery_margin;comment:交易所交割保证金;"`                                  // 交易所交割保证金
	ReserveBalance                 decimal.Decimal `json:"reserve_balance" gorm:"column:reserve_balance;comment:保底期货结算准备金;"`                                                   // 保底期货结算准备金
	CurrencyID                     string          `json:"currency_id" gorm:"column:currency_id;comment:币种代码;"`                                                                // 币种代码
	PreFundMortgageIn              decimal.Decimal `json:"pre_fund_mortgage_in" gorm:"column:pre_fund_mortgage_in;comment:上次货币质入金额;"`                                          // 上次货币质入金额
	PreFundMortgageOut             decimal.Decimal `json:"pre_fund_mortgage_out" gorm:"column:pre_fund_mortgage_out;comment:上次货币质出金额;"`                                        // 上次货币质出金额
	FundMortgageIn                 decimal.Decimal `json:"fund_mortgage_in" gorm:"column:fund_mortgage_in;comment:货币质入金额;"`                                                    // 货币质入金额
	FundMortgageOut                decimal.Decimal `json:"fund_mortgage_out" gorm:"column:fund_mortgage_out;comment:货币质出金额;"`                                                  // 货币质出金额
	FundMortgageAvailable          decimal.Decimal `json:"fund_mortgage_available" gorm:"column:fund_mortgage_available;comment:货币质押余额;"`                                      // 货币质押余额
	MortgageableFund               decimal.Decimal `json:"mortgageable_fund" gorm:"column:mortgageable_fund;comment:可质押货币金额;"`                                                 // 可质押货币金额
	SpecProductMargin              decimal.Decimal `json:"spec_product_margin" gorm:"column:spec_product_margin;comment:特殊产品占用保证金;"`                                           // 特殊产品占用保证金
	SpecProductFrozenMargin        decimal.Decimal `json:"spec_product_frozen_margin" gorm:"column:spec_product_frozen_margin;comment:特殊产品冻结保证金;"`                             // 特殊产品冻结保证金
	SpecProductCommission          decimal.Decimal `json:"spec_product_commission" gorm:"column:spec_product_commission;comment:特殊产品手续费;"`                                     // 特殊产品手续费
	SpecProductFrozenCommission    decimal.Decimal `json:"spec_product_frozen_commission" gorm:"column:spec_product_frozen_commission;comment:特殊产品冻结手续费;"`                     // 特殊产品冻结手续费
	SpecProductPositionProfit      decimal.Decimal `json:"spec_product_position_profit" gorm:"column:spec_product_position_profit;comment:特殊产品持仓盈亏;"`                          // 特殊产品持仓盈亏
	SpecProductCloseProfit         decimal.Decimal `json:"spec_product_close_profit" gorm:"column:spec_product_close_profit;comment:特殊产品平仓盈亏;"`                                // 特殊产品平仓盈亏
	SpecProductPositionProfitByAlg decimal.Decimal `json:"spec_product_position_profit_by_alg" gorm:"column:spec_product_position_profit_by_alg;comment:根据持仓盈亏算法计算的特殊产品持仓盈亏;"` // 根据持仓盈亏算法计算的特殊产品持仓盈亏
	SpecProductExchangeMargin      decimal.Decimal `json:"spec_product_exchange_margin" gorm:"column:spec_product_exchange_margin;comment:特殊产品交易所保证金;"`                        // 特殊产品交易所保证金
	BizType                        int32           `json:"biz_type" gorm:"column:biz_type;comment:业务类型;"`                                                                      // 业务类型
	FrozenSwap                     decimal.Decimal `json:"frozen_swap" gorm:"column:frozen_swap;comment:延时换汇冻结金额;"`                                                            // 延时换汇冻结金额
	RemainSwap                     decimal.Decimal `json:"remain_swap" gorm:"column:remain_swap;comment:剩余换汇额度;"`                                                              // 剩余换汇额度
}

type PosDateType int32

const (
	//THOST_FTDC_PSD_Today   PosDateType = '0' //纯今日
	THOST_FTDC_PSD_Today   PosDateType = '1' //今日(49)
	THOST_FTDC_PSD_History PosDateType = '2' //旧仓(50)
)

// 持仓
type Position struct {
	BrokerID           string          `json:"broker_id" gorm:"column:broker_id;comment:经纪公司代码;index:idx_broker_id,type:btree"`
	InstrumentID       string          `json:"instrument_id" gorm:"column:instrument_id;comment:合约代码;index:idx_instrument_id,type:btree"`
	InstrumentName     string          `json:"instrument_name" gorm:"column:instrument_name;comment:合约名;"`
	ExchangeID         string          `json:"exchange_id" gorm:"column:exchange_id;comment:交易所代码;"`
	InvestorID         string          `json:"investor_id" gorm:"column:investor_id;comment:投资者代码;index:idx_investor_id,type:btree"`
	PosiDirection      int32           `json:"posi_direction" gorm:"column:posi_direction;comment:持仓多空方向;"`
	HedgeFlag          int32           `json:"hedge_flag" gorm:"column:hedge_flag;comment:投机套保标志;"`
	PositionDate       int32           `json:"position_date" gorm:"column:position_date;comment:持仓日期;index:idx_position_date,type:btree"`
	Position           int             `json:"position" gorm:"column:position;comment:当前总持仓;"`
	YdPosition         int             `json:"yd_position" gorm:"column:yd_position;comment:上日持仓;"`
	TodayPosition      int             `json:"today_position" gorm:"column:today_position;comment:表示今新开仓;"`
	OpenVolume         int             `json:"open_volume" gorm:"column:open_volume;comment:开仓量;"`
	CloseVolume        int             `json:"close_volume" gorm:"column:close_volume;comment:平仓量;"`
	OpenAmount         decimal.Decimal `json:"open_amount" gorm:"column:open_amount;comment:开仓金额;"`
	CloseAmount        decimal.Decimal `json:"close_amount" gorm:"column:close_amount;comment:平仓金额;"`
	PositionCost       decimal.Decimal `json:"position_cost" gorm:"column:position_cost;comment:持仓成本=上日持仓*昨结算价*合约乘数+SUM(今日持仓*开仓价*合约乘数);"`
	CloseProfit        decimal.Decimal `json:"close_profit" gorm:"column:close_profit;comment:平仓盈亏;"`
	PositionProfit     decimal.Decimal `json:"position_profit" gorm:"column:position_profit;comment:持仓盈亏;"`
	OpenCost           decimal.Decimal `json:"open_cost" gorm:"column:open_cost;comment:开仓成本;"`
	LongFrozen         int             `json:"long_frozen" gorm:"column:long_frozen;comment:多头冻结;"`
	ShortFrozen        int             `json:"short_frozen" gorm:"column:short_frozen;comment:空头冻结;"`
	LongFrozenAmount   decimal.Decimal `json:"long_frozen_amount" gorm:"column:long_frozen_amount;comment:开仓冻结金额;"`
	ShortFrozenAmount  decimal.Decimal `json:"short_frozen_amount" gorm:"column:short_frozen_amount;comment:开仓冻结金额;"`
	PreMargin          decimal.Decimal `json:"pre_margin" gorm:"column:pre_margin;comment:上次占用的保证金;"`
	UseMargin          decimal.Decimal `json:"use_margin" gorm:"column:use_margin;comment:占用的保证金;"`
	FrozenMargin       decimal.Decimal `json:"frozen_margin" gorm:"column:frozen_margin;comment:冻结的保证金;"`
	FrozenCash         decimal.Decimal `json:"frozen_cash" gorm:"column:frozen_cash;comment:冻结的资金;"`
	FrozenCommission   decimal.Decimal `json:"frozen_commission" gorm:"column:frozen_commission;comment:冻结的手续费;"`
	CashIn             decimal.Decimal `json:"cash_in" gorm:"column:cash_in;comment:资金差额;"`
	Commission         decimal.Decimal `json:"commission" gorm:"column:commission;comment:手续费;"`
	PreSettlementPrice decimal.Decimal `json:"pre_settlement_price" gorm:"column:pre_settlement_price;comment:上次结算价;"`
	SettlementPrice    decimal.Decimal `json:"settlement_price" gorm:"column:settlement_price;comment:本次结算价;"`
	SettlementID       int             `json:"settlement_id" gorm:"column:settlement_id;comment:结算编号;index:idx_settlement_id,type:btree"`
	ExchangeMargin     decimal.Decimal `json:"exchange_margin" gorm:"column:exchange_margin;comment:交易所保证金;"`
	CombPosition       int             `json:"comb_position" gorm:"column:comb_position;comment:组合成交形成的持仓;"`
	CombLongFrozen     int             `json:"comb_long_frozen" gorm:"column:comb_long_frozen;comment:组合多头冻结;"`
	CombShortFrozen    int             `json:"comb_short_frozen" gorm:"column:comb_short_frozen;comment:组合空头冻结;"`
	CloseProfitByDate  decimal.Decimal `json:"close_profit_by_date" gorm:"column:close_profit_by_date;comment:逐日盯市平仓盈亏;"`
	CloseProfitByTrade decimal.Decimal `json:"close_profit_by_trade" gorm:"column:close_profit_by_trade;comment:逐笔对冲平仓盈亏;"`
	StrikeFrozen       int             `json:"strike_frozen" gorm:"column:strike_frozen;comment:执行冻结;"`
	StrikeFrozenAmount decimal.Decimal `json:"strike_frozen_amount" gorm:"column:strike_frozen_amount;comment:执行冻结金额;"`
	AbandonFrozen      int             `json:"abandon_frozen" gorm:"column:abandon_frozen;comment:放弃执行冻结;"`
	YdStrikeFrozen     int             `json:"yd_strike_frozen" gorm:"column:yd_strike_frozen;comment:执行冻结的昨仓;"`
	PositionCostOffset decimal.Decimal `json:"position_cost_offset" gorm:"column:position_cost_offset;comment:大商所持仓成本差值，只有大商所使用;"`
	Reserve1           string          `json:"reserve1" gorm:"column:reserve1;comment:保留的无效字段;"`
}

// vm:合约数量乘数 x 持仓数量
func (impl *Position) AvgPrice(vm int64) decimal.Decimal {
	if vm == 0 {
		return decimal.Zero
	}
	return impl.OpenCost.Div(decimal.NewFromInt(vm * int64(impl.Position)))
}
