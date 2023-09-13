package gormwrapper

type GormAssociationInterface interface {
	Append(...interface{}) error
	Clear() error
	Count() int64
	Delete(...interface{}) error
	Find(interface{}, ...interface{}) error
	Replace(...interface{}) error
}
