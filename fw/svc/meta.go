package svc

const Key_Meta = "tk:meta"

type Meta interface {
	Get(key string) interface{}
	Load(val interface{}) error
}
