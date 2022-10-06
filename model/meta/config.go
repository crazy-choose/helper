package meta

import "time"

type Postgres struct {
	Host         string `json:"host" toml:"host" mapstructure:"host"`                         // 服务器地址
	Port         string `json:"port" toml:"port" mapstructure:"port"`                         // 服务器端口
	DbName       string `json:"dbName" toml:"dbName" mapstructure:"dbName"`                   // 数据库名
	Username     string `json:"userName" toml:"userName" mapstructure:"userName"`             // 数据库用户名
	Password     string `json:"passWord" toml:"passWord" mapstructure:"passWord"`             // 数据库密码
	MaxIdleConns int    `json:"maxIdleConns" toml:"maxIdleConns" mapstructure:"maxIdleConns"` // 空闲中的最大连接数
	MaxOpenConns int    `json:"MaxOpenConns" toml:"MaxOpenConns" mapstructure:"MaxOpenConns"` // 打开到数据库的最大连接数
	LogMode      bool   `json:"logMode" toml:"logMode" mapstructure:"logMode"`                // 是否开启Gorm全局日志
	LogZap       bool   `json:"logZap" toml:"logZap" mapstructure:"logZap"`                   // 是否通过zap写入日志文件
	LogLevel     string `json:"logLevel" toml:"logLevel" mapstructure:"logLevel"`             // 日志级别
}

type Redis struct {
	Enabled      bool
	Host         string
	AuthKey      string
	DbIndex      int
	PoolSize     int
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Zap struct {
	Name          string `mapstructure:"name" json:"name" yaml:"name"`                              // 文jianming
	Level         string `mapstructure:"level" json:"level" yaml:"level"`                           // 级别
	Format        string `mapstructure:"format" json:"format" yaml:"format"`                        // 输出
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                        // 日志前缀
	Director      string `mapstructure:"director" json:"director"  yaml:"director"`                 // 日志文件夹
	LinkName      string `mapstructure:"link-name" json:"linkName" yaml:"link-name"`                // 软链接名称
	ShowLine      bool   `mapstructure:"show-line" json:"showLine" yaml:"showLine"`                 // 显示行
	EncodeLevel   string `mapstructure:"encode-level" json:"encodeLevel" yaml:"encode-level"`       // 编码级
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktraceKey" yaml:"stacktrace-key"` // 栈名
	LogInConsole  bool   `mapstructure:"log-in-console" json:"logInConsole" yaml:"log-in-console"`  // 输出控制台
}

type CTPDeploy struct {
	TradeFront []string
	QuoteFront []string
	BrokerID   string
	Password   string
	UserID     string
	AppID      string
	AuthCode   string
	FundPassWD string
	DelivDay   int
	Products   []string
}
