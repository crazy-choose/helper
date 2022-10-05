package db

import (
	"database/sql"
	"errors"
	"github.com/crazy-choose/helper/log"
	"github.com/crazy-choose/helper/model/meta"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"math"
)

var (
	option           meta.PG
	instance         *Session
	ErrUninitialized = errors.New("pg client not initialized")
)

func verifyConfig(option meta.PG) {
	if option.DbName == "" {
		log.Panic("dbname is empty")
	}
}

func genDSN(option meta.PG) string {
	return "host =" + option.Host + " user=" + option.Username + " password=" + option.Password + " dbname=" + option.DbName + " port=" + option.Port + " sslmode=disable TimeZone=Asia/Shanghai"
}

type Session struct {
	*gorm.DB
}

func (s *Session) Transaction(fc func(tx *Session) error, opts ...*sql.TxOptions) error {
	return s.DB.Transaction(func(tx *gorm.DB) error {
		return fc(&Session{tx})
	}, opts...)
}

func InitPg(opt meta.PG) {
	verifyConfig(opt)
	pgConfig := postgres.Config{
		PreferSimpleProtocol: true,
		DSN:                  genDSN(option), // DSN data source name
	}
	if db, err := gorm.Open(postgres.New(pgConfig)); err != nil {
		log.Error(err.Error())
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(option.MaxIdleConns)
		sqlDB.SetMaxOpenConns(option.MaxOpenConns)
		instance = &Session{
			db,
		}
		log.Info("InitPg success")
	}
}

func Conn() *Session {
	if instance == nil {
		InitPg(option)
	}
	return instance
}

func Transaction(fc func(tx *Session) error, opts ...*sql.TxOptions) (err error) {
	return Conn().DB.Transaction(func(tx *gorm.DB) error {
		return fc(&Session{tx})
	}, opts...)
}

type Pagination struct {
	Limit      int         `json:"limit,omitempty;query:limit"`
	Page       int         `json:"page,omitempty;query:page"`
	Sort       string      `json:"sort,omitempty;query:sort"`
	TotalRows  int64       `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
	Rows       interface{} `json:"rows"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "Id desc"
	}
	return p.Sort
}

func Paginate(pagination *Pagination, txn *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	txn.Count(&totalRows)
	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}
