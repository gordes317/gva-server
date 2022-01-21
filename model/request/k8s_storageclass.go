package request

import "gin-vue-admin/model"

type StorageClassSearch struct {
	model.StorageClass
	PageInfo
}

type StorageClassSearchUser struct {
	StorageClassSearch
	model.User
}
