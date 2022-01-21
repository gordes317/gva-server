package request

import "gin-vue-admin/model"

type PVSearch struct {
	model.PV
	PageInfo
}

type PVSearchUser struct {
	PVSearch
	model.User
}
