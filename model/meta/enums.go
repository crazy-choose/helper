package meta

type DirectionType byte // 买卖方向类型
const (
	DirectionBuy  DirectionType = '0' // 买
	DirectionSell DirectionType = '1' // 卖
	DirectionNone DirectionType = 'x'
)

type PosDirectionType byte // 买卖方向类型
const (
	THOST_FTDC_PD_Net   PosDirectionType = '1' //
	THOST_FTDC_PD_Long  PosDirectionType = '2' //多
	THOST_FTDC_PD_Short PosDirectionType = '3' //空
)

type OffsetFlagType byte // 开平标志类型
const (
	THOST_FTDC_OF_Open            OffsetFlagType = '0' // 开仓
	THOST_FTDC_OF_Close           OffsetFlagType = '1' // 平仓
	THOST_FTDC_OF_ForceClose      OffsetFlagType = '2' // 强平
	THOST_FTDC_OF_CloseToday      OffsetFlagType = '3' // 平今
	THOST_FTDC_OF_CloseYesterday  OffsetFlagType = '4' // 平昨
	THOST_FTDC_OF_ForceOff        OffsetFlagType = '5' // 强减
	THOST_FTDC_OF_LocalForceClose OffsetFlagType = '6' // 本地强平
)

// 订单报价类型
type OrderPriceType byte

const (
	THOST_FTDC_OPT_AnyPrice                OrderPriceType = '1' // 市价单
	THOST_FTDC_OPT_LimitPrice              OrderPriceType = '2' // 限价单
	THOST_FTDC_OPT_BestPrice               byte           = '3' // 限价单
	THOST_FTDC_OPT_LastPrice               byte           = '4' // 限价单
	THOST_FTDC_OPT_LastPricePlusOneTicks   byte           = '5' // 限价单
	THOST_FTDC_OPT_LastPricePlusTwoTicks   byte           = '6' // 限价单
	THOST_FTDC_OPT_LastPricePlusThreeTicks byte           = '7' // 限价单
	THOST_FTDC_OPT_AskPrice1               byte           = '8' // 限价单
	THOST_FTDC_OPT_AskPrice1PlusOneTicks   byte           = '9' // 限价单
	THOST_FTDC_OPT_AskPrice1PlusTwoTicks   byte           = 'A' // 限价单
	THOST_FTDC_OPT_AskPrice1PlusThreeTicks byte           = 'B' // 限价单
	THOST_FTDC_OPT_BidPrice1               byte           = 'C' // 限价单
	THOST_FTDC_OPT_BidPrice1PlusOneTicks   byte           = 'D' // 限价单
	THOST_FTDC_OPT_BidPrice1PlusTwoTicks   byte           = 'E' // 限价单
	THOST_FTDC_OPT_BidPrice1PlusThreeTicks byte           = 'F' // 限价单
	THOST_FTDC_OPT_FiveLevelPrice          byte           = 'G' // 限价单
)

//有效期类型
type TimeConditionType byte

const (
	THOST_FTDC_TC_IOC byte = '1' //立即完成，否则撤销
	THOST_FTDC_TC_GFS byte = '2'
	THOST_FTDC_TC_GFD byte = '3' //当日有效
	THOST_FTDC_TC_GTD byte = '4'
	THOST_FTDC_TC_GTC byte = '5'
	THOST_FTDC_TC_GFA byte = '6'
)

type OrderStatusType byte // 报单状态类型
const (
	OrderStatusAllTraded             OrderStatusType = '0' // 全部成交
	OrderStatusPartTradedQueueing    OrderStatusType = '1' // 部分成交还在队列中
	OrderStatusPartTradedNotQueueing OrderStatusType = '2' // 部分成交不在队列中
	OrderStatusNoTradeQueueing       OrderStatusType = '3' // 未成交还在队列中
	OrderStatusNoTradeNotQueueing    OrderStatusType = '4' // 未成交不在队列中
	OrderStatusCanceled              OrderStatusType = '5' // 撤单
	OrderStatusUnknown               OrderStatusType = 'a' // 未知
	OrderStatusNotTouched            OrderStatusType = 'b' // 尚未触发
	OrderStatusTouched               OrderStatusType = 'c' // 已触发
)

type ActionFlagType byte

const (
	THOST_FTDC_AF_Delete ActionFlagType = '0'
	THOST_FTDC_AF_Modify ActionFlagType = '3'
)

type InstrumentStatusType byte

const (
	THOST_FTDC_IS_BeforeTrading   InstrumentStatusType = '0'
	THOST_FTDC_IS_NoTrading       InstrumentStatusType = '1'
	THOST_FTDC_IS_Continous       InstrumentStatusType = '2'
	THOST_FTDC_IS_AuctionOrdering InstrumentStatusType = '3'
	THOST_FTDC_IS_AuctionBalance  InstrumentStatusType = '4'
	THOST_FTDC_IS_AuctionMatch    InstrumentStatusType = '5'
	THOST_FTDC_IS_Closed          InstrumentStatusType = '6'
)

type InstStatusEnterReasonType byte

const (
	THOST_FTDC_IER_Automatic InstStatusEnterReasonType = '1'
	THOST_FTDC_IER_Manual    InstStatusEnterReasonType = '2'
	THOST_FTDC_IER_Fuse      InstStatusEnterReasonType = '3'
)



type HedgeFlagTypes byte // 投机套保标志类型
const (
	HedgeFlagSpeculation HedgeFlagTypes = '1' // 投机
	HedgeFlagArbitrage   HedgeFlagTypes = '2' // 套利
	HedgeFlagHedge       HedgeFlagTypes = '3' // 套保
	HedgeFlagMarketMaker HedgeFlagTypes = '5' // 做市商
	HedgeFlagSpecHedge   HedgeFlagTypes = '6' // 第一腿投机第二腿套保 大商所专用
	HedgeFlagHedgeSpec   HedgeFlagTypes = '7' // 第一腿套保第二腿投机  大商所专用
)

type ProductClassTypes byte // 产品类型类型
const (
	ProductClassFutures     ProductClassTypes = '1' // 期货
	ProductClassOptions     ProductClassTypes = '2' // 期货期权
	ProductClassCombination ProductClassTypes = '3' // 组合
	ProductClassSpot        ProductClassTypes = '4' // 即期
	ProductClassEFP         ProductClassTypes = '5' // 期转现
	ProductClassSpotOption  ProductClassTypes = '6' // 现货期权
)

type OptionsTypes byte // 期权类型
const (
	OptionsTypeCallOptions OptionsTypes = '1' // 看涨
	OptionsTypePutOptions  OptionsTypes = '2' // 看跌
)

type CombinationTypes byte // 组合类型
const (
	CombinationTypeFuture CombinationTypes = '0' // 期货组合
	CombinationTypeBUL    CombinationTypes = '1' // 垂直价差BUL
	CombinationTypeBER    CombinationTypes = '2' // 垂直价差BER
	CombinationTypeSTD    CombinationTypes = '3' // 跨式组合
	CombinationTypeSTG    CombinationTypes = '4' // 宽跨式组合
	CombinationTypePRT    CombinationTypes = '5' // 备兑组合
	CombinationTypeCLD    CombinationTypes = '6' // 时间价差组合
)
