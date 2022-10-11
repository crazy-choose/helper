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
