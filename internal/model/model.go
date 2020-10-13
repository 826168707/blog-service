package model

import (
	"blog-service/global"
	"blog-service/pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
	ID  		uint32	`gorm:"primary_key" json:"id"`
	CreatedBy	string	`json:"created_by"`
	ModifiedBy	string	`json:"modified_by"`
	CreatedOn	uint32	`json:"created_on"`
	ModifiedOn	uint32	`json:"modified_on"`
	DeletedOn	uint32	`json:"deleted_on"`
	IsDel		uint8	`json:"is_del"`
}

// 创建DB连接引擎
func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB,error) {
	db,err := gorm.Open(databaseSetting.DBType,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
			databaseSetting.Username,
			databaseSetting.Password,
			databaseSetting.Host,
			databaseSetting.DBName,
			databaseSetting.Charset,
			databaseSetting.ParseTime,
		))
	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}

	db.SingularTable(true)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)

	return db,nil
}