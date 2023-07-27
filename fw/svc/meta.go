package svc

const Key_Meta = "tk:meta"

type Meta interface {
	Load(key string, val interface{}) error
}
