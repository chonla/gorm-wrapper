package gormwrapper

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GormInterface interface {
	Association(string) GormAssociationInterface
	AutoMigrate(...interface{}) error
	Begin(...*sql.TxOptions) GormInterface
	Clauses(...clause.Expression) GormInterface
	Commit() GormInterface
	Count(*int64) GormInterface
	Create(interface{}) GormInterface
	Debug() GormInterface
	Delete(interface{}, ...interface{}) GormInterface
	Distinct(...interface{}) GormInterface
	Error() error
	Exec(string, ...interface{}) GormInterface
	Find(interface{}, ...interface{}) GormInterface
	First(interface{}, ...interface{}) GormInterface
	Group(string) GormInterface
	Joins(string, ...interface{}) GormInterface
	Last(interface{}, ...interface{}) GormInterface
	Limit(int) GormInterface
	Model(interface{}) GormInterface
	Offset(int) GormInterface
	Order(interface{}) GormInterface
	Pluck(string, interface{}) GormInterface
	Preload(string, ...interface{}) GormInterface
	Raw(string, ...interface{}) GormInterface
	Rollback() GormInterface
	Save(interface{}) GormInterface
	Scan(interface{}) GormInterface
	Scopes(...func(*gorm.DB) *gorm.DB) GormInterface
	Select(string, ...interface{}) GormInterface
	Session(*gorm.Session) GormInterface
	SetupJoinTable(model interface{}, field string, joinTable interface{}) error
	Statement() *gorm.Statement
	Unscoped() GormInterface
	Update(string, interface{}) GormInterface
	Updates(interface{}) GormInterface
	UpdateColumns(interface{}) GormInterface
	Where(interface{}, ...interface{}) GormInterface
	WithContext(context.Context) GormInterface
}
