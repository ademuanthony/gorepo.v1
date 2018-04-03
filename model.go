package gorepo

type IModel interface {
	GetId() interface{}

	Validate() error
}

