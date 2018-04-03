package gorepo

import "github.com/jinzhu/gorm"


type IRepository interface {

	// Insert puts a new instance of the give IModel in the database
	Insert(model IModel) (interface{}, error)

	Update(model IModel) (error)

	Save(model IModel) (interface{}, error)

	FindById(receiver IModel, uint interface{}) (error)

	FindFirst(receiver IModel, where string, args ...interface{}) (error)

	FindAll(models interface{}, where string, args ...interface{}) (err error)

	FindAllWithLimit(models interface{}, offset int, limit int, where string, args ...interface{}) (err error)

	Count(models interface{}, where string, args ...interface{}) (result int, err error)

	Delete(model IModel, where string, args ...interface{}) error

	// NewRecord check if the model exist in the store
	NewRecord(model IModel) bool

	RawSql(sql string, value ...interface{}) (db *gorm.DB)
}
