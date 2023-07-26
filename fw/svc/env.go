package svc

const Key_Env = "tk:env"

type Env interface {
	AppEnv() string
}
