package request

import "gin-vue-admin/model"

type PVCSearch struct {
	model.PVC
	PageInfo
}

type PVCSearchUser struct {
	PVCSearch
	model.User
}
