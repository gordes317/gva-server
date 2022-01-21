package request

import "gin-vue-admin/model"

type ConfigMapSearch struct {
	model.ConfigMap
	PageInfo
}

type ConfigMapSearchUser struct {
	ConfigMapSearch
	model.User
}
