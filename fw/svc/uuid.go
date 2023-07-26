package svc

const Key_UUID = "tk:uuid"

type UUID interface {
	NewID() string
}
