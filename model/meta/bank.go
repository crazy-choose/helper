package meta

import "github.com/shopspring/decimal"

/*
	TradeCode：202001 期货发起银行资金转期货
	TradeCode：202002 期货发起期货资金转银行
	BankID：返回银行在期货公司的代码
	BankBranchID：填0000
	BrokerBranchID：一般为空
	CustomerName：客户姓名
	BankPassWord：返回银行的密码，都是“*”密码不显示
	Password：返回期货账户的密码，都是“*”密码不显示
	TransferStatus: 转账交易的最后状态
	TradeAmount：转账金额
	pRspInfo：响应信息
*/

type ReqBankFutureTransferInfo struct {
	TradeCode         string          // 业务功能码
	BankID            string          // 银行代码
	BankBranchID      string          // 银行分支机构代码
	BrokerID          string          // 期商代码
	BrokerBranchID    string          // 期商分支机构代码
	TradeDate         string          // 交易日期
	TradeTime         string          // 交易时间
	BankSerial        string          // 银行流水号
	TradingDay        string          // 交易系统日期
	PlateSerial       int             // 银期平台消息流水号
	LastFragment      string          // 最后分片标志
	SessionID         int             // 会话号
	CustomerName      string          // 客户姓名
	IdCardType        string          // 证件类型
	IdentifiedCardNo  string          // 证件号码
	CustType          string          // 客户类型
	BankAccount       string          // 银行帐号
	BankPassWord      string          // 银行密码
	AccountID         string          // 投资者帐号
	Password          string          // 期货密码
	InstallID         int             // 安装编号
	FutureSerial      int             // 期货公司流水号
	UserID            string          // 用户标识
	VerifyCertNoFlag  string          // 验证客户证件号码标志
	CurrencyID        string          // 币种代码
	TradeAmount       decimal.Decimal // 转帐金额
	FutureFetchAmount decimal.Decimal // 期货可取金额
	FeePayFlag        string          // 费用支付标志
	CustFee           decimal.Decimal // 应收客户费用
	BrokerFee         decimal.Decimal // 应收期货公司费用
	Message           string          // 发送方给接收方的消息
	Digest            string          // 摘要
	BankAccType       string          // 银行帐号类型
	DeviceID          string          // 渠道标志
	BankSecuAccType   string          // 期货单位帐号类型
	BrokerIDByBank    string          // 期货公司银行编码
	BankSecuAcc       string          // 期货单位帐号
	BankPwdFlag       string          // 银行密码标志
	SecuPwdFlag       string          // 期货资金密码核对标志
	OperNo            string          // 交易柜员
	RequestID         int             // 请求编号
	TID               int             // 交易ID
	TransferStatus    string          // 转账交易状态
	LongCustomerName  string          // 长客户姓名
}

// RspBankFutureTransferInfo 银期转账响应
type RspBankFutureTransferInfo struct {
	TradeCode         string          // 业务功能码
	BankID            string          // 银行代码
	BankBranchID      string          // 银行分支机构代码
	BrokerID          string          // 期商代码
	BrokerBranchID    string          // 期商分支机构代码
	TradeDate         string          // 交易日期
	TradeTime         string          // 交易时间
	BankSerial        string          // 银行流水号
	TradingDay        string          // 交易系统日期
	PlateSerial       int             // 银期平台消息流水号
	LastFragment      string          // 最后分片标志
	SessionID         int             // 会话号
	CustomerName      string          // 客户姓名
	IdCardType        string          // 证件类型
	IdentifiedCardNo  string          // 证件号码
	CustType          string          // 客户类型
	BankAccount       string          // 银行帐号
	BankPassWord      string          // 银行密码
	AccountID         string          // 投资者帐号
	Password          string          // 期货密码
	InstallID         int             // 安装编号
	FutureSerial      int             // 期货公司流水号
	UserID            string          // 用户标识
	VerifyCertNoFlag  string          // 验证客户证件号码标志
	CurrencyID        string          // 币种代码
	TradeAmount       decimal.Decimal // 转帐金额
	FutureFetchAmount decimal.Decimal // 期货可取金额
	FeePayFlag        string          // 费用支付标志
	CustFee           decimal.Decimal // 应收客户费用
	BrokerFee         decimal.Decimal // 应收期货公司费用
	Message           string          // 发送方给接收方的消息
	Digest            string          // 摘要
	BankAccType       string          // 银行帐号类型
	DeviceID          string          // 渠道标志
	BankSecuAccType   string          // 期货单位帐号类型
	BrokerIDByBank    string          // 期货公司银行编码
	BankSecuAcc       string          // 期货单位帐号
	BankPwdFlag       string          // 银行密码标志
	SecuPwdFlag       string          // 期货资金密码核对标志
	OperNo            string          // 交易柜员
	RequestID         int             // 请求编号
	TID               int             // 交易ID
	TransferStatus    string          // 转账交易状态
	LongCustomerName  string          // 长客户姓名
}
