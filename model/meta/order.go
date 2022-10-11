package meta

import "github.com/shopspring/decimal"

// 订单信息 && 报单响应
type OrderInfo struct {
	BrokerID             string          `json:"broker_id" gorm:"column:broker_id;comment:经纪公司代码;index:idx_broker_id,type:btree"`                // 经纪公司代码
	InvestorID           string          `json:"investor_id" gorm:"column:investor_id;comment:投资者代码;index:idx_investor_id,type:btree"`           // 投资者代码
	OrderRef             string          `json:"order_ref" gorm:"column:order_ref;comment:报单引用;index:idx_order_ref,type:btree"`                  // 报单引用
	UserID               string          `json:"user_id" gorm:"column:user_id;comment:用户代码;index:idx_user_id,type:btree"`                        // 用户代码
	OrderPriceType       string          `json:"order_price_type" gorm:"column:order_price_type;comment:报单价格条件;"`                                // 报单价格条件
	Direction            byte            `json:"direction" gorm:"column:direction;comment:买卖方向;"`                                                // 买卖方向
	CombOffsetFlag       string          `json:"comb_offset_flag" gorm:"column:comb_offset_flag;comment:组合开平标志;"`                                // 组合开平标志
	CombHedgeFlag        string          `json:"comb_hedge_flag" gorm:"column:comb_hedge_flag;comment:组合投机套保标志;"`                                // 组合投机套保标志
	LimitPrice           decimal.Decimal `json:"limit_price" gorm:"column:limit_price;comment:价格;"`                                              // 价格
	VolumeTotalOriginal  int             `json:"volume_total_original" gorm:"column:volume_total_original;comment:数量;"`                          // 数量
	TimeCondition        string          `json:"time_condition" gorm:"column:time_condition;comment:有效期类型;"`                                     // 有效期类型
	GTDDate              string          `json:"gtd_date" gorm:"column:gtd_date;comment:GTD日期;"`                                                 // GTD日期
	VolumeCondition      string          `json:"volume_condition" gorm:"column:volume_condition;comment:成交量类型;"`                                 // 成交量类型
	MinVolume            int             `json:"min_volume" gorm:"column:min_volume;comment:最小成交量;"`                                             // 最小成交量
	ContingentCondition  string          `json:"contingent_condition" gorm:"column:contingent_condition;comment:触发条件;"`                          // 触发条件
	StopPrice            decimal.Decimal `json:"stop_price" gorm:"column:stop_price;comment:止损价;"`                                               // 止损价
	ForceCloseReason     byte            `json:"force_close_reason" gorm:"column:force_close_reason;comment:强平原因;"`                              // 强平原因
	IsAutoSuspend        int             `json:"is_auto_suspend" gorm:"column:is_auto_suspend;comment:自动挂起标志;"`                                  // 自动挂起标志
	BusinessUnit         string          `json:"business_unit" gorm:"column:business_unit;comment:业务单元;"`                                        // 业务单元
	RequestID            int             `json:"request_id" gorm:"column:request_id;comment:请求编号;index:idx_request_id,type:btree"`               // 请求编号
	OrderLocalID         string          `json:"order_local_id" gorm:"column:order_local_id;comment:本地报单编号;index:idx_order_local_id,type:btree"` // 本地报单编号
	ExchangeID           string          `json:"exchange_id" gorm:"column:exchange_id;comment:交易所代码;"`                                           // 交易所代码
	ParticipantID        string          `json:"participant_id" gorm:"column:participant_id;comment:会员代码;"`                                      // 会员代码
	ClientID             string          `json:"client_id" gorm:"column:client_id;comment:客户代码;"`                                                // 客户代码
	AccountID            string          `json:"account_id" gorm:"column:account_id;comment:资金账号;"`                                              // 资金账号
	InstrumentID         string          `json:"instrument_id" gorm:"column:instrument_id;comment:合约代码;index:idx_instrument_id,type:btree"`      // 合约代码
	ExchangeInstID       string          `json:"exchange_inst_id" gorm:"column:exchange_inst_id;comment:合约在交易所的代码;"`                             // 合约在交易所的代码
	TraderID             string          `json:"trader_id" gorm:"column:trader_id;comment:交易所交易员代码;"`                                            // 交易所交易员代码
	InstallID            int             `json:"install_id" gorm:"column:install_id;comment:安装编号;"`                                              // 安装编号
	OrderSubmitStatus    string          `json:"order_submit_status" gorm:"column:order_submit_status;comment:报单提交状态;"`                          // 报单提交状态
	NotifySequence       int             `json:"notify_sequence" gorm:"column:notify_sequence;comment:报单提示序号;"`                                  // 报单提示序号
	TradingDay           string          `json:"trading_day" gorm:"column:trading_day;comment:交易日;"`                                             // 交易日
	SettlementID         int             `json:"settlement_id" gorm:"column:settlement_id;comment:结算编号;"`                                        // 结算编号
	OrderSysID           string          `json:"order_sys_id" gorm:"column:order_sys_id;comment:报单编号;index:idx_order_sys_id,type:btree"`         // 报单编号
	OrderSource          string          `json:"order_source" gorm:"column:order_source;comment:报单来源;index:idx_order_source,type:btree"`         // 报单来源
	OrderStatus          byte            `json:"order_status" gorm:"column:order_status;comment:报单状态;"`                                          // 报单状态
	OrderType            string          `json:"order_type" gorm:"column:order_type;comment:报单类型;"`                                              // 报单类型
	VolumeTraded         int             `json:"volume_traded" gorm:"column:volume_traded;comment:今成交数量;"`                                       // 今成交数量
	VolumeTotal          int             `json:"volume_total" gorm:"column:volume_total;comment:剩余数量;"`                                          // 剩余数量
	InsertDate           string          `json:"insert_date" gorm:"column:insert_date;comment:报单日期;index:idx_insert_date,type:btree"`            // 报单日期
	InsertTime           string          `json:"insert_time" gorm:"column:insert_time;comment:委托时间;"`                                            // 委托时间
	ActiveTime           string          `json:"active_time" gorm:"column:active_time;comment:激活时间;"`                                            // 激活时间
	SuspendTime          string          `json:"suspend_time" gorm:"column:suspend_time;comment:挂起时间;"`                                          // 挂起时间
	UpdateTime           string          `json:"update_time" gorm:"column:update_time;comment:最后修改时间;"`                                          // 最后修改时间
	CancelTime           string          `json:"cancel_time" gorm:"column:cancel_time;comment:撤销时间;"`                                            // 撤销时间
	ActiveTraderID       string          `json:"active_trader_id" gorm:"column:active_trader_id;comment:最后修改交易所交易员代码;"`                          // 最后修改交易所交易员代码
	ClearingPartID       string          `json:"clearing_part_id" gorm:"column:clearing_part_id;comment:结算会员编号;"`                                // 结算会员编号
	SequenceNo           int             `json:"sequence_no" gorm:"column:sequence_no;comment:序号;"`                                              // 序号
	FrontID              int             `json:"front_id" gorm:"column:front_id;comment:前置编号;"`                                                  // 前置编号
	SessionID            int             `json:"session_id" gorm:"column:session_id;comment:会话编号;"`                                              // 会话编号
	UserProductInfo      string          `json:"user_product_info" gorm:"column:user_product_info;comment:用户端产品信息;"`                             // 用户端产品信息
	StatusMsg            string          `json:"status_msg" gorm:"column:status_msg;comment:状态信息;"`                                              // 状态信息
	UserForceClose       int             `json:"user_force_close" gorm:"column:user_force_close;comment:用户强评标志;"`                                // 用户强评标志
	ActiveUserID         string          `json:"active_user_id" gorm:"column:active_user_id;comment:操作用户代码;"`                                    // 操作用户代码
	BrokerOrderSeq       int             `json:"broker_order_seq" gorm:"column:broker_order_seq;comment:经纪公司报单编号;"`                              // 经纪公司报单编号
	RelativeOrderSysID   string          `json:"relative_order_sys_id" gorm:"column:relative_order_sys_id;comment:相关报单;"`                        // 相关报单
	ZCETotalTradedVolume int             `json:"zce_total_traded_volume" gorm:"column:zce_total_traded_volume;comment:郑商所成交数量;"`                 // 郑商所成交数量
	IsSwapOrder          int             `json:"is_swap_order" gorm:"column:is_swap_order;comment:互换单标志;"`                                       // 互换单标志
	BranchID             string          `json:"branch_id" gorm:"column:branch_id;comment:营业部编号;"`                                               // 营业部编号
	InvestUnitID         string          `json:"invest_unit_id" gorm:"column:invest_unit_id;comment:投资单元代码;"`                                    // 投资单元代码
	CurrencyID           string          `json:"currency_id" gorm:"column:currency_id;comment:币种代码;"`                                            // 币种代码
	MacAddress           string          `json:"mac_address" gorm:"column:mac_address;comment:Mac地址;"`                                           // Mac地址
	IPAddress            string          `json:"ip_address" gorm:"column:ip_address;comment:IP地址;"`                                              // IP地址
	reserve1             string          `json:"reserve_1" gorm:"column:reserve_1;comment:保留的无效字段1;"`                                            // 保留的无效字段
	reserve2             string          `json:"reserve_2" gorm:"column:reserve_2;comment:保留的无效字段2;"`                                            // 保留的无效字段
	reserve3             string          `json:"reserve_3" gorm:"column:reserve_3;comment:保留的无效字段3;"`                                            // 保留的无效字段
	//Error                RspInfo         `json:"error"`                   //errorInfo(add)
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
	ActionFlag        byte            `json:"action_flag"`         // 操作标志
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
