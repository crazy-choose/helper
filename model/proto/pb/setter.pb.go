package pb

import (
	"unsafe"
)

func (x *TickInfo) Swigcptr() uintptr {
	return uintptr(unsafe.Pointer(x))
}

func (x *TickInfo) SwigIsCThostFtdcInstrumentCommissionRateField() {}
func (x *TickInfo) SetInstrumentID(value string)                   { x.InstrumentID = value }
func (x *TickInfo) SetExchangeID(value string) {
	x.ExchangeID = value
}
func (x *TickInfo) SetExchangeInstID(value string) {
	x.ExchangeInstID = value
}
func (x *TickInfo) SetLastPrice(value float64) {
	x.LastPrice = value
}
func (x *TickInfo) SetPreSettlementPrice(value float64) {
	x.PreSettlementPrice = value
}
func (x *TickInfo) SetPreClosePrice(value float64) {
	x.PreClosePrice = value
}
func (x *TickInfo) SetPreOpenInterest(value float64) {
	x.PreOpenInterest = value
}
func (x *TickInfo) SetOpenPrice(value float64) {
	x.OpenPrice = value
}
func (x *TickInfo) SetClosePrice(value float64) {
	x.ClosePrice = value
}
func (x *TickInfo) SetSettlementPrice(value float64) {
	x.SettlementPrice = value
}
func (x *TickInfo) SetHighestPrice(value float64) {
	x.HighestPrice = value
}
func (x *TickInfo) SetLowestPrice(value float64) {
	x.LowestPrice = value
}
func (x *TickInfo) SetVolume(value int) {
	x.Volume = int32(value)
}
func (x *TickInfo) SetTurnover(value float64) {
	x.Turnover = value
}
func (x *TickInfo) SetOpenInterest(value float64) {
	x.OpenInterest = value
}
func (x *TickInfo) SetUpdateTime(value string) {
	x.UpdateTime = value
}
func (x *TickInfo) SetUpdateMillisec(value int) {
	x.UpdateMillisec = int32(value)
}
func (x *TickInfo) SetAveragePrice(value float64) {
	x.AveragePrice = value
}
func (x *TickInfo) SetActionDay(value string) {
	x.ActionDay = value
}
func (x *TickInfo) SetUpperLimitPrice(value float64) {
	x.UpperLimitPrice = value
}
func (x *TickInfo) SetLowerLimitPrice(value float64) {
	x.LowerLimitPrice = value
}
func (x *TickInfo) SetBidPrice1(value float64) {
	x.BidPrice1 = value
}
func (x *TickInfo) SetBidVolume1(value int) {
	x.BidVolume1 = int32(value)
}

func (x *TickInfo) SetAskPrice1(value float64) {
	x.AskPrice1 = value
}

func (x *TickInfo) SetAskVolume1(value int) {
	x.AskVolume1 = int32(value)
}

func (x *TickInfo) SetBidPrice2(value float64) {
	x.BidPrice2 = value
}

func (x *TickInfo) SetBidVolume2(value int) {
	x.BidVolume2 = int32(value)
}

func (x *TickInfo) SetAskPrice2(value float64) {
	x.AskPrice2 = value
}

func (x *TickInfo) SetAskVolume2(value int) {
	x.AskVolume2 = int32(value)
}

func (x *TickInfo) SetBidPrice3(value float64) {
	x.BidPrice3 = value
}

func (x *TickInfo) SetBidVolume3(value int) {
	x.BidVolume3 = int32(value)
}

func (x *TickInfo) SetAskPrice3(value float64) {
	x.AskPrice3 = value
}

func (x *TickInfo) SetAskVolume3(value int) {
	x.AskVolume3 = int32(value)
}

func (x *TickInfo) SetBidPrice4(value float64) {
	x.BidPrice4 = value
}

func (x *TickInfo) SetBidVolume4(value int) {
	x.BidVolume4 = int32(value)
}

func (x *TickInfo) SetAskPrice4(value float64) {
	x.AskPrice4 = value
}

func (x *TickInfo) SetAskVolume4(value int) {
	x.AskVolume4 = int32(value)
}

func (x *TickInfo) SetBidPrice5(value float64) {
	x.BidPrice5 = value
}

func (x *TickInfo) SetBidVolume5(value int) {
	x.BidVolume5 = int32(value)
}

func (x *TickInfo) SetAskPrice5(value float64) {
	x.AskPrice5 = value
}

func (x *TickInfo) SetAskVolume5(value int) {
	x.AskVolume5 = int32(value)
}

func (x *TickInfo) SetPreDelta(value float64) {
	x.PreDelta = value
}

func (x *TickInfo) SetCurrDelta(value float64) {
	x.CurrDelta = value
}

func (x *TickInfo) SetBandingUpperPrice(value float64) {
	x.BandingUpperPrice = value
}

func (x *TickInfo) SetBandingLowerPrice(value float64) {
	x.BandingLowerPrice = value
}

func (x *TickInfo) SetActionDirection(value int32) {
	x.ActionDirection = value
}

func (x *TickInfo) SetSyncDeltaSequenceNo(value int32) {
	x.SyncDeltaSequenceNo = value
}

func (x *TickInfo) SetAnalyze(value string) {
	x.Analyze = value
}

// AccountInfo setters
func (a *AccountInfo) Swigcptr() uintptr                        { return uintptr(unsafe.Pointer(a)) }
func (a *AccountInfo) SwigIsCThostFtdcTradingAccountField()     {}
func (a *AccountInfo) SetBrokerID(value string)                 { a.BrokerID = value }
func (a *AccountInfo) SetAccountID(value string)                { a.AccountID = value }
func (a *AccountInfo) SetPreMortgage(value float64)             { a.PreMortgage = value }
func (a *AccountInfo) SetPreCredit(value float64)               { a.PreCredit = value }
func (a *AccountInfo) SetPreDeposit(value float64)              { a.PreDeposit = value }
func (a *AccountInfo) SetPreBalance(value float64)              { a.PreBalance = value }
func (a *AccountInfo) SetPreMargin(value float64)               { a.PreMargin = value }
func (a *AccountInfo) SetInterestBase(value float64)            { a.InterestBase = value }
func (a *AccountInfo) SetInterest(value float64)                { a.Interest = value }
func (a *AccountInfo) SetDeposit(value float64)                 { a.Deposit = value }
func (a *AccountInfo) SetWithdraw(value float64)                { a.Withdraw = value }
func (a *AccountInfo) SetFrozenMargin(value float64)            { a.FrozenMargin = value }
func (a *AccountInfo) SetFrozenCash(value float64)              { a.FrozenCash = value }
func (a *AccountInfo) SetFrozenCommission(value float64)        { a.FrozenCommission = value }
func (a *AccountInfo) SetCurrMargin(value float64)              { a.CurrMargin = value }
func (a *AccountInfo) SetCashIn(value float64)                  { a.CashIn = value }
func (a *AccountInfo) SetCommission(value float64)              { a.Commission = value }
func (a *AccountInfo) SetCloseProfit(value float64)             { a.CloseProfit = value }
func (a *AccountInfo) SetPositionProfit(value float64)          { a.PositionProfit = value }
func (a *AccountInfo) SetBalance(value float64)                 { a.Balance = value }
func (a *AccountInfo) SetAvailable(value float64)               { a.Available = value }
func (a *AccountInfo) SetWithdrawQuota(value float64)           { a.WithdrawQuota = value }
func (a *AccountInfo) SetReserve(value float64)                 { a.Reserve = value }
func (a *AccountInfo) SetTradingDay(value string)               { a.TradingDay = value }
func (a *AccountInfo) SetSettlementID(value int32)              { a.SettlementID = value }
func (a *AccountInfo) SetCredit(value float64)                  { a.Credit = value }
func (a *AccountInfo) SetMortgage(value float64)                { a.Mortgage = value }
func (a *AccountInfo) SetExchangeMargin(value float64)          { a.ExchangeMargin = value }
func (a *AccountInfo) SetDeliveryMargin(value float64)          { a.DeliveryMargin = value }
func (a *AccountInfo) SetExchangeDeliveryMargin(value float64)  { a.ExchangeDeliveryMargin = value }
func (a *AccountInfo) SetReserveBalance(value float64)          { a.ReserveBalance = value }
func (a *AccountInfo) SetCurrencyID(value string)               { a.CurrencyID = value }
func (a *AccountInfo) SetPreFundMortgageIn(value float64)       { a.PreFundMortgageIn = value }
func (a *AccountInfo) SetPreFundMortgageOut(value float64)      { a.PreFundMortgageOut = value }
func (a *AccountInfo) SetFundMortgageIn(value float64)          { a.FundMortgageIn = value }
func (a *AccountInfo) SetFundMortgageOut(value float64)         { a.FundMortgageOut = value }
func (a *AccountInfo) SetFundMortgageAvailable(value float64)   { a.FundMortgageAvailable = value }
func (a *AccountInfo) SetMortgageableFund(value float64)        { a.MortgageableFund = value }
func (a *AccountInfo) SetSpecProductMargin(value float64)       { a.SpecProductMargin = value }
func (a *AccountInfo) SetSpecProductFrozenMargin(value float64) { a.SpecProductFrozenMargin = value }
func (a *AccountInfo) SetSpecProductCommission(value float64)   { a.SpecProductCommission = value }
func (a *AccountInfo) SetSpecProductFrozenCommission(value float64) {
	a.SpecProductFrozenCommission = value
}
func (a *AccountInfo) SetSpecProductPositionProfit(value float64) {
	a.SpecProductPositionProfit = value
}
func (a *AccountInfo) SetSpecProductCloseProfit(value float64) { a.SpecProductCloseProfit = value }
func (a *AccountInfo) SetSpecProductPositionProfitByAlg(value float64) {
	a.SpecProductPositionProfitByAlg = value
}
func (a *AccountInfo) SetSpecProductExchangeMargin(value float64) {
	a.SpecProductExchangeMargin = value
}
func (a *AccountInfo) SetBizType(value int32)      { a.BizType = value }
func (a *AccountInfo) SetFrozenSwap(value float64) { a.FrozenSwap = value }
func (a *AccountInfo) SetRemainSwap(value float64) { a.RemainSwap = value }

// Position setters
func (a *Position) Swigcptr() uintptr                      { return uintptr(unsafe.Pointer(a)) }
func (a *Position) SwigIsCThostFtdcInvestorPositionField() {}
func (p *Position) SetBrokerID(value string)               { p.BrokerID = value }
func (p *Position) SetInstrumentID(value string)           { p.InstrumentID = value }
func (p *Position) SetInstrumentName(value string)         { p.InstrumentName = value }
func (p *Position) SetExchangeID(value string)             { p.ExchangeID = value }
func (p *Position) SetInvestorID(value string)             { p.InvestorID = value }
func (p *Position) SetPosiDirection(value int32)           { p.PosiDirection = value }
func (p *Position) SetHedgeFlag(value int32)               { p.HedgeFlag = value }
func (p *Position) SetPositionDate(value PosDateType)      { p.PositionDate = value }
func (p *Position) SetPosition(value int32)                { p.Position = value }
func (p *Position) SetYdPosition(value int32)              { p.YdPosition = value }
func (p *Position) SetTodayPosition(value int32)           { p.TodayPosition = value }
func (p *Position) SetOpenVolume(value int32)              { p.OpenVolume = value }
func (p *Position) SetCloseVolume(value int32)             { p.CloseVolume = value }
func (p *Position) SetOpenAmount(value float64)            { p.OpenAmount = value }
func (p *Position) SetCloseAmount(value float64)           { p.CloseAmount = value }
func (p *Position) SetPositionCost(value float64)          { p.PositionCost = value }
func (p *Position) SetCloseProfit(value float64)           { p.CloseProfit = value }
func (p *Position) SetPositionProfit(value float64)        { p.PositionProfit = value }
func (p *Position) SetOpenCost(value float64)              { p.OpenCost = value }
func (p *Position) SetLongFrozen(value int32)              { p.LongFrozen = value }
func (p *Position) SetShortFrozen(value int32)             { p.ShortFrozen = value }
func (p *Position) SetLongFrozenAmount(value float64)      { p.LongFrozenAmount = value }
func (p *Position) SetShortFrozenAmount(value float64)     { p.ShortFrozenAmount = value }
func (p *Position) SetPreMargin(value float64)             { p.PreMargin = value }
func (p *Position) SetUseMargin(value float64)             { p.UseMargin = value }
func (p *Position) SetFrozenMargin(value float64)          { p.FrozenMargin = value }
func (p *Position) SetFrozenCash(value float64)            { p.FrozenCash = value }
func (p *Position) SetFrozenCommission(value float64)      { p.FrozenCommission = value }
func (p *Position) SetCashIn(value float64)                { p.CashIn = value }
func (p *Position) SetCommission(value float64)            { p.Commission = value }
func (p *Position) SetPreSettlementPrice(value float64)    { p.PreSettlementPrice = value }
func (p *Position) SetSettlementPrice(value float64)       { p.SettlementPrice = value }
func (p *Position) SetSettlementID(value int32)            { p.SettlementID = value }
func (p *Position) SetExchangeMargin(value float64)        { p.ExchangeMargin = value }
func (p *Position) SetCombPosition(value int32)            { p.CombPosition = value }
func (p *Position) SetCombLongFrozen(value int32)          { p.CombLongFrozen = value }
func (p *Position) SetCombShortFrozen(value int32)         { p.CombShortFrozen = value }
func (p *Position) SetCloseProfitByDate(value float64)     { p.CloseProfitByDate = value }
func (p *Position) SetCloseProfitByTrade(value float64)    { p.CloseProfitByTrade = value }
func (p *Position) SetStrikeFrozen(value int32)            { p.StrikeFrozen = value }
func (p *Position) SetStrikeFrozenAmount(value float64)    { p.StrikeFrozenAmount = value }
func (p *Position) SetAbandonFrozen(value int32)           { p.AbandonFrozen = value }
func (p *Position) SetYdStrikeFrozen(value int32)          { p.YdStrikeFrozen = value }
func (p *Position) SetPositionCostOffset(value float64)    { p.PositionCostOffset = value }
func (p *Position) SetReserve1(value string)               { p.Reserve1 = value }

// OrderInfo setters
func (a *OrderInfo) Swigcptr() uintptr                   { return uintptr(unsafe.Pointer(a)) }
func (a *OrderInfo) SwigIsCThostFtdcOrderField()         {}
func (o *OrderInfo) SetBrokerID(value string)            { o.BrokerID = value }
func (o *OrderInfo) SetInvestorID(value string)          { o.InvestorID = value }
func (o *OrderInfo) SetOrderRef(value string)            { o.OrderRef = value }
func (o *OrderInfo) SetUserID(value string)              { o.UserID = value }
func (o *OrderInfo) SetOrderPriceType(value int32)       { o.OrderPriceType = value }
func (o *OrderInfo) SetDirection(value int32)            { o.Direction = value }
func (o *OrderInfo) SetCombOffsetFlag(value string)      { o.CombOffsetFlag = value }
func (o *OrderInfo) SetCombHedgeFlag(value string)       { o.CombHedgeFlag = value }
func (o *OrderInfo) SetLimitPrice(value float64)         { o.LimitPrice = value }
func (o *OrderInfo) SetVolumeTotalOriginal(value int32)  { o.VolumeTotalOriginal = value }
func (o *OrderInfo) SetTimeCondition(value string)       { o.TimeCondition = value }
func (o *OrderInfo) SetGTDDate(value string)             { o.GTDDate = value }
func (o *OrderInfo) SetVolumeCondition(value string)     { o.VolumeCondition = value }
func (o *OrderInfo) SetMinVolume(value int32)            { o.MinVolume = value }
func (o *OrderInfo) SetContingentCondition(value string) { o.ContingentCondition = value }
func (o *OrderInfo) SetStopPrice(value float64)          { o.StopPrice = value }
func (o *OrderInfo) SetForceCloseReason(value int32)     { o.ForceCloseReason = value }
func (o *OrderInfo) SetIsAutoSuspend(value int32)        { o.IsAutoSuspend = value }
func (o *OrderInfo) SetBusinessUnit(value string)        { o.BusinessUnit = value }
func (o *OrderInfo) SetRequestID(value int32)            { o.RequestID = value }
func (o *OrderInfo) SetOrderLocalID(value string)        { o.OrderLocalID = value }
func (o *OrderInfo) SetExchangeID(value string)          { o.ExchangeID = value }
func (o *OrderInfo) SetParticipantID(value string)       { o.ParticipantID = value }
func (o *OrderInfo) SetClientID(value string)            { o.ClientID = value }
func (o *OrderInfo) SetAccountID(value string)           { o.AccountID = value }
func (o *OrderInfo) SetInstrumentID(value string)        { o.InstrumentID = value }
func (o *OrderInfo) SetExchangeInstID(value string)      { o.ExchangeInstID = value }
func (o *OrderInfo) SetTraderID(value string)            { o.TraderID = value }
func (o *OrderInfo) SetInstallID(value int32)            { o.InstallID = value }
func (o *OrderInfo) SetOrderSubmitStatus(value string)   { o.OrderSubmitStatus = value }
func (o *OrderInfo) SetNotifySequence(value int32)       { o.NotifySequence = value }
func (o *OrderInfo) SetTradingDay(value string)          { o.TradingDay = value }
func (o *OrderInfo) SetSettlementID(value int32)         { o.SettlementID = value }
func (o *OrderInfo) SetOrderSysID(value string)          { o.OrderSysID = value }
func (o *OrderInfo) SetOrderSource(value string)         { o.OrderSource = value }
func (o *OrderInfo) SetOrderStatus(value int32)          { o.OrderStatus = value }
func (o *OrderInfo) SetOrderType(value string)           { o.OrderType = value }
func (o *OrderInfo) SetVolumeTraded(value int32)         { o.VolumeTraded = value }
func (o *OrderInfo) SetVolumeTotal(value int32)          { o.VolumeTotal = value }
func (o *OrderInfo) SetInsertDate(value string)          { o.InsertDate = value }
func (o *OrderInfo) SetInsertTime(value string)          { o.InsertTime = value }
func (o *OrderInfo) SetActiveTime(value string)          { o.ActiveTime = value }
func (o *OrderInfo) SetSuspendTime(value string)         { o.SuspendTime = value }
func (o *OrderInfo) SetUpdateTime(value string)          { o.UpdateTime = value }
func (o *OrderInfo) SetCancelTime(value string)          { o.CancelTime = value }
func (o *OrderInfo) SetActiveTraderID(value string)      { o.ActiveTraderID = value }
func (o *OrderInfo) SetClearingPartID(value string)      { o.ClearingPartID = value }
func (o *OrderInfo) SetSequenceNo(value int32)           { o.SequenceNo = value }
func (o *OrderInfo) SetFrontID(value int32)              { o.FrontID = value }
func (o *OrderInfo) SetSessionID(value int32)            { o.SessionID = value }
func (o *OrderInfo) SetUserProductInfo(value string)     { o.UserProductInfo = value }
func (o *OrderInfo) SetStatusMsg(value string)           { o.StatusMsg = value }
func (o *OrderInfo) SetUserForceClose(value int32)       { o.UserForceClose = value }
func (o *OrderInfo) SetActiveUserID(value string)        { o.ActiveUserID = value }
func (o *OrderInfo) SetBrokerOrderSeq(value int32)       { o.BrokerOrderSeq = value }
func (o *OrderInfo) SetRelativeOrderSysID(value string)  { o.RelativeOrderSysID = value }
func (o *OrderInfo) SetZceTotalTradedVolume(value int32) { o.ZceTotalTradedVolume = value }
func (o *OrderInfo) SetIsSwapOrder(value int32)          { o.IsSwapOrder = value }
func (o *OrderInfo) SetBranchID(value string)            { o.BranchID = value }
func (o *OrderInfo) SetInvestUnitID(value string)        { o.InvestUnitID = value }
func (o *OrderInfo) SetCurrencyID(value string)          { o.CurrencyID = value }
func (o *OrderInfo) SetMacAddress(value string)          { o.MacAddress = value }
func (o *OrderInfo) SetIPAddress(value string)           { o.IPAddress = value }
func (o *OrderInfo) SetReserve1(value string)            { o.Reserve1 = value }
func (o *OrderInfo) SetReserve2(value string)            { o.Reserve2 = value }
func (o *OrderInfo) SetReserve3(value string)            { o.Reserve3 = value }

// OrderActionInfo setters
func (a *OrderActionInfo) Swigcptr() uintptr                 { return uintptr(unsafe.Pointer(a)) }
func (a *OrderActionInfo) SwigIsCThostFtdcOrderActionField() {}
func (o *OrderActionInfo) SetBrokerID(value string)          { o.BrokerID = value }
func (o *OrderActionInfo) SetInvestorID(value string)        { o.InvestorID = value }
func (o *OrderActionInfo) SetOrderActionRef(value int32)     { o.OrderActionRef = value }
func (o *OrderActionInfo) SetOrderRef(value string)          { o.OrderRef = value }
func (o *OrderActionInfo) SetRequestID(value int32)          { o.RequestID = value }
func (o *OrderActionInfo) SetFrontID(value int32)            { o.FrontID = value }
func (o *OrderActionInfo) SetSessionID(value int32)          { o.SessionID = value }
func (o *OrderActionInfo) SetExchangeID(value string)        { o.ExchangeID = value }
func (o *OrderActionInfo) SetOrderSysID(value string)        { o.OrderSysID = value }
func (o *OrderActionInfo) SetActionFlag(value int32)         { o.ActionFlag = value }
func (o *OrderActionInfo) SetLimitPrice(value float64)       { o.LimitPrice = value }
func (o *OrderActionInfo) SetVolumeChange(value int32)       { o.VolumeChange = value }
func (o *OrderActionInfo) SetActionDate(value string)        { o.ActionDate = value }
func (o *OrderActionInfo) SetActionTime(value string)        { o.ActionTime = value }
func (o *OrderActionInfo) SetTraderID(value string)          { o.TraderID = value }
func (o *OrderActionInfo) SetInstallID(value int32)          { o.InstallID = value }
func (o *OrderActionInfo) SetOrderLocalID(value string)      { o.OrderLocalID = value }
func (o *OrderActionInfo) SetActionLocalID(value string)     { o.ActionLocalID = value }
func (o *OrderActionInfo) SetParticipantID(value string)     { o.ParticipantID = value }
func (o *OrderActionInfo) SetClientID(value string)          { o.ClientID = value }
func (o *OrderActionInfo) SetBusinessUnit(value string)      { o.BusinessUnit = value }
func (o *OrderActionInfo) SetOrderActionStatus(value string) { o.OrderActionStatus = value }
func (o *OrderActionInfo) SetUserID(value string)            { o.UserID = value }
func (o *OrderActionInfo) SetStatusMsg(value string)         { o.StatusMsg = value }
func (o *OrderActionInfo) SetReserve1(value string)          { o.Reserve1 = value }
func (o *OrderActionInfo) SetBranchID(value string)          { o.BranchID = value }
func (o *OrderActionInfo) SetInvestUnitID(value string)      { o.InvestUnitID = value }
func (o *OrderActionInfo) SetReserve2(value string)          { o.Reserve2 = value }
func (o *OrderActionInfo) SetMacAddress(value string)        { o.MacAddress = value }
func (o *OrderActionInfo) SetInstrumentID(value string)      { o.InstrumentID = value }
func (o *OrderActionInfo) SetIPAddress(value string)         { o.IPAddress = value }

// TradeInfo setters
func (a *TradeInfo) Swigcptr() uintptr              { return uintptr(unsafe.Pointer(a)) }
func (a *TradeInfo) SwigIsCThostFtdcTradeField()    {}
func (t *TradeInfo) SetBrokerID(value string)       { t.BrokerID = value }
func (t *TradeInfo) SetInvestorID(value string)     { t.InvestorID = value }
func (t *TradeInfo) SetReserve1(value string)       { t.Reserve1 = value }
func (t *TradeInfo) SetOrderRef(value string)       { t.OrderRef = value }
func (t *TradeInfo) SetUserID(value string)         { t.UserID = value }
func (t *TradeInfo) SetExchangeID(value string)     { t.ExchangeID = value }
func (t *TradeInfo) SetTradeID(value string)        { t.TradeID = value }
func (t *TradeInfo) SetDirection(value int32)       { t.Direction = value }
func (t *TradeInfo) SetOrderSysID(value string)     { t.OrderSysID = value }
func (t *TradeInfo) SetParticipantID(value string)  { t.ParticipantID = value }
func (t *TradeInfo) SetClientID(value string)       { t.ClientID = value }
func (t *TradeInfo) SetTradingRole(value string)    { t.TradingRole = value }
func (t *TradeInfo) SetReserve2(value string)       { t.Reserve2 = value }
func (t *TradeInfo) SetOffsetFlag(value int32)      { t.OffsetFlag = value }
func (t *TradeInfo) SetHedgeFlag(value string)      { t.HedgeFlag = value }
func (t *TradeInfo) SetPrice(value float64)         { t.Price = value }
func (t *TradeInfo) SetVolume(value int32)          { t.Volume = value }
func (t *TradeInfo) SetTradeDate(value string)      { t.TradeDate = value }
func (t *TradeInfo) SetTradeTime(value string)      { t.TradeTime = value }
func (t *TradeInfo) SetTradeType(value int32)       { t.TradeType = value }
func (t *TradeInfo) SetPriceSource(value string)    { t.PriceSource = value }
func (t *TradeInfo) SetTraderID(value string)       { t.TraderID = value }
func (t *TradeInfo) SetOrderLocalID(value string)   { t.OrderLocalID = value }
func (t *TradeInfo) SetClearingPartID(value string) { t.ClearingPartID = value }
func (t *TradeInfo) SetBusinessUnit(value string)   { t.BusinessUnit = value }
func (t *TradeInfo) SetSequenceNo(value int32)      { t.SequenceNo = value }
func (t *TradeInfo) SetTradingDay(value string)     { t.TradingDay = value }
func (t *TradeInfo) SetSettlementID(value int32)    { t.SettlementID = value }
func (t *TradeInfo) SetBrokerOrderSeq(value int32)  { t.BrokerOrderSeq = value }
func (t *TradeInfo) SetTradeSource(value int32)     { t.TradeSource = value }
func (t *TradeInfo) SetInvestUnitID(value string)   { t.InvestUnitID = value }
func (t *TradeInfo) SetInstrumentID(value string)   { t.InstrumentID = value }
func (t *TradeInfo) SetExchangeInstID(value string) { t.ExchangeInstID = value }

// BankFutureTransferInfo setters
func (a *BankFutureTransferInfo) Swigcptr() uintptr                  { return uintptr(unsafe.Pointer(a)) }
func (a *BankFutureTransferInfo) SwigIsCThostFtdcRspTransferField()  {}
func (b *BankFutureTransferInfo) SetTradeCode(value string)          { b.TradeCode = value }
func (b *BankFutureTransferInfo) SetBankID(value string)             { b.BankID = value }
func (b *BankFutureTransferInfo) SetBankBranchID(value string)       { b.BankBranchID = value }
func (b *BankFutureTransferInfo) SetBrokerID(value string)           { b.BrokerID = value }
func (b *BankFutureTransferInfo) SetBrokerBranchID(value string)     { b.BrokerBranchID = value }
func (b *BankFutureTransferInfo) SetTradeDate(value string)          { b.TradeDate = value }
func (b *BankFutureTransferInfo) SetTradeTime(value string)          { b.TradeTime = value }
func (b *BankFutureTransferInfo) SetBankSerial(value string)         { b.BankSerial = value }
func (b *BankFutureTransferInfo) SetTradingDay(value string)         { b.TradingDay = value }
func (b *BankFutureTransferInfo) SetPlateSerial(value int32)         { b.PlateSerial = value }
func (b *BankFutureTransferInfo) SetLastFragment(value string)       { b.LastFragment = value }
func (b *BankFutureTransferInfo) SetSessionID(value int32)           { b.SessionID = value }
func (b *BankFutureTransferInfo) SetCustomerName(value string)       { b.CustomerName = value }
func (b *BankFutureTransferInfo) SetIdCardType(value string)         { b.IdCardType = value }
func (b *BankFutureTransferInfo) SetIdentifiedCardNo(value string)   { b.IdentifiedCardNo = value }
func (b *BankFutureTransferInfo) SetCustType(value string)           { b.CustType = value }
func (b *BankFutureTransferInfo) SetBankAccount(value string)        { b.BankAccount = value }
func (b *BankFutureTransferInfo) SetBankPassWord(value string)       { b.BankPassWord = value }
func (b *BankFutureTransferInfo) SetAccountID(value string)          { b.AccountID = value }
func (b *BankFutureTransferInfo) SetPassword(value string)           { b.Password = value }
func (b *BankFutureTransferInfo) SetInstallID(value int32)           { b.InstallID = value }
func (b *BankFutureTransferInfo) SetFutureSerial(value int32)        { b.FutureSerial = value }
func (b *BankFutureTransferInfo) SetUserID(value string)             { b.UserID = value }
func (b *BankFutureTransferInfo) SetVerifyCertNoFlag(value string)   { b.VerifyCertNoFlag = value }
func (b *BankFutureTransferInfo) SetCurrencyID(value string)         { b.CurrencyID = value }
func (b *BankFutureTransferInfo) SetTradeAmount(value float64)       { b.TradeAmount = value }
func (b *BankFutureTransferInfo) SetFutureFetchAmount(value float64) { b.FutureFetchAmount = value }
func (b *BankFutureTransferInfo) SetFeePayFlag(value string)         { b.FeePayFlag = value }
func (b *BankFutureTransferInfo) SetCustFee(value float64)           { b.CustFee = value }
func (b *BankFutureTransferInfo) SetBrokerFee(value float64)         { b.BrokerFee = value }
func (b *BankFutureTransferInfo) SetMessage(value string)            { b.Message = value }
func (b *BankFutureTransferInfo) SetDigest(value string)             { b.Digest = value }
func (b *BankFutureTransferInfo) SetBankAccType(value string)        { b.BankAccType = value }
func (b *BankFutureTransferInfo) SetDeviceID(value string)           { b.DeviceID = value }
func (b *BankFutureTransferInfo) SetBankSecuAccType(value string)    { b.BankSecuAccType = value }
func (b *BankFutureTransferInfo) SetBrokerIDByBank(value string)     { b.BrokerIDByBank = value }
func (b *BankFutureTransferInfo) SetBankSecuAcc(value string)        { b.BankSecuAcc = value }
func (b *BankFutureTransferInfo) SetBankPwdFlag(value string)        { b.BankPwdFlag = value }
func (b *BankFutureTransferInfo) SetSecuPwdFlag(value string)        { b.SecuPwdFlag = value }
func (b *BankFutureTransferInfo) SetOperNo(value string)             { b.OperNo = value }
func (b *BankFutureTransferInfo) SetRequestID(value int32)           { b.RequestID = value }
func (b *BankFutureTransferInfo) SetTID(value int32)                 { b.TID = value }
func (b *BankFutureTransferInfo) SetTransferStatus(value string)     { b.TransferStatus = value }
func (b *BankFutureTransferInfo) SetErrorID(value int)               { b.ErrorID = int32(value) }
func (b *BankFutureTransferInfo) SetErrorMsg(value string)           { b.ErrorMsg = value }
func (b *BankFutureTransferInfo) SetLongCustomerName(value string)   { b.LongCustomerName = value }
