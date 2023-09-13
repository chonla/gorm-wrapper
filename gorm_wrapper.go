package gormwrapper

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GormDB struct {
	db *gorm.DB
}

func Open(dialector gorm.Dialector, opts ...gorm.Option) (GormInterface, error) {
	db, err := gorm.Open(dialector, opts...)
	return Wrap(db), err
}

func Wrap(db *gorm.DB) GormInterface {
	return &GormDB{
		db: db,
	}
}

func (o *GormDB) Association(column string) GormAssociationInterface {
	return WrapAssociation(o.db.Association(column))
}

func (o *GormDB) AutoMigrate(dst ...interface{}) error {
	return o.db.AutoMigrate(dst...)
}

func (o *GormDB) Begin(opts ...*sql.TxOptions) GormInterface {
	return Wrap(o.db.Begin(opts...))
}

func (o *GormDB) Clauses(conds ...clause.Expression) GormInterface {
	return Wrap(o.db.Clauses(conds...))
}

func (o *GormDB) Commit() GormInterface {
	return Wrap(o.db.Commit())
}

func (o *GormDB) Count(count *int64) GormInterface {
	return Wrap(o.db.Count(count))
}

func (o *GormDB) Create(value interface{}) GormInterface {
	return Wrap(o.db.Create(value))
}

func (o *GormDB) Debug() GormInterface {
	return Wrap(o.db.Debug())
}

func (o *GormDB) Delete(value interface{}, conds ...interface{}) GormInterface {
	return Wrap(o.db.Delete(value, conds...))
}

func (o *GormDB) Distinct(args ...interface{}) GormInterface {
	return Wrap(o.db.Distinct(args...))
}

func (o *GormDB) Error() error {
	return o.db.Error
}

func (o *GormDB) Exec(sql string, values ...interface{}) GormInterface {
	return Wrap(o.db.Exec(sql, values...))
}

func (o *GormDB) Find(dest interface{}, conds ...interface{}) GormInterface {
	return Wrap(o.db.Find(dest, conds...))
}

func (o *GormDB) First(dest interface{}, conds ...interface{}) GormInterface {
	return Wrap(o.db.First(dest, conds...))
}

func (o *GormDB) Group(name string) GormInterface {
	return Wrap(o.db.Group(name))
}

func (o *GormDB) Joins(query string, conds ...interface{}) GormInterface {
	return Wrap(o.db.Joins(query, conds...))
}

func (o *GormDB) Last(dest interface{}, conds ...interface{}) GormInterface {
	return Wrap(o.db.Last(dest, conds...))
}

func (o *GormDB) Limit(limit int) GormInterface {
	return Wrap(o.db.Limit(limit))
}

func (o *GormDB) Model(value interface{}) GormInterface {
	return Wrap(o.db.Model(value))
}

func (o *GormDB) Offset(offset int) GormInterface {
	return Wrap(o.db.Offset(offset))
}

func (o *GormDB) Order(value interface{}) GormInterface {
	return Wrap(o.db.Order(value))
}

func (o *GormDB) Preload(query string, args ...interface{}) GormInterface {
	return Wrap(o.db.Preload(query, args...))
}

func (o *GormDB) Pluck(column string, dest interface{}) GormInterface {
	return Wrap(o.db.Pluck(column, dest))
}

func (o *GormDB) Raw(query string, args ...interface{}) GormInterface {
	return Wrap(o.db.Raw(query, args...))
}

func (o *GormDB) Rollback() GormInterface {
	return Wrap(o.db.Rollback())
}

func (o *GormDB) Save(value interface{}) GormInterface {
	return Wrap(o.db.Save(value))
}

func (o *GormDB) Scan(dest interface{}) GormInterface {
	return Wrap(o.db.Scan(dest))
}

func (o *GormDB) Scopes(funcs ...func(*gorm.DB) *gorm.DB) GormInterface {
	return Wrap(o.db.Scopes(funcs...))
}

func (o *GormDB) Select(query string, args ...interface{}) GormInterface {
	return Wrap(o.db.Select(query, args...))
}

func (o *GormDB) Session(config *gorm.Session) GormInterface {
	return Wrap(o.db.Session(config))
}

func (o *GormDB) SetupJoinTable(model interface{}, field string, joinTable interface{}) error {
	return o.db.SetupJoinTable(model, field, joinTable)
}

func (o *GormDB) Statement() *gorm.Statement {
	return o.db.Statement
}

func (o *GormDB) Update(column string, value interface{}) GormInterface {
	return Wrap(o.db.Update(column, value))
}

func (o *GormDB) Updates(values interface{}) GormInterface {
	return Wrap(o.db.Updates(values))
}

func (o *GormDB) UpdateColumns(values interface{}) GormInterface {
	return Wrap(o.db.UpdateColumns(values))
}

func (o *GormDB) Unscoped() GormInterface {
	return Wrap(o.db.Unscoped())
}

func (o *GormDB) Where(query interface{}, args ...interface{}) GormInterface {
	return Wrap(o.db.Where(query, args...))
}

func (o *GormDB) WithContext(ctx context.Context) GormInterface {
	return Wrap(o.db.WithContext(ctx))
}
