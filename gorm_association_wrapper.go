package gormwrapper

import (
	"gorm.io/gorm"
)

type GormAssociate struct {
	assoc *gorm.Association
}

func WrapAssociation(assoc *gorm.Association) GormAssociationInterface {
	return &GormAssociate{
		assoc: assoc,
	}
}

func (o *GormAssociate) Append(values ...interface{}) error {
	return o.assoc.Append(values)
}

func (o *GormAssociate) Clear() error {
	return o.assoc.Clear()
}

func (o *GormAssociate) Count() int64 {
	return o.assoc.Count()
}

func (o *GormAssociate) Find(out interface{}, conds ...interface{}) error {
	return o.assoc.Find(out, conds)
}

func (o *GormAssociate) Replace(values ...interface{}) error {
	return o.assoc.Replace(values)
}

func (o *GormAssociate) Delete(values ...interface{}) error {
	return o.assoc.Delete(values)
}
