package request

import (
	"gin-vue-admin/model"
)

type IngressSearch struct {
	model.Ingress
	PageInfo
}

type IngressSearchUser struct {
	IngressSearch
	model.User
}
