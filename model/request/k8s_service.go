package request

import "gin-vue-admin/model"

type ServiceSearch struct {
	model.Service
	PageInfo
}

type ServiceSearchUser struct {
	ServiceSearch
	model.User
}
