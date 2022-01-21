package request

import "gin-vue-admin/model"

type SelectNamespacesSearch struct {
	model.SelectNamespaces
	PageInfo
}

type SelectNamespacesSearchUser struct {
	model.User
	SelectNamespacesSearch
}
type PageInfoUser struct {
	model.User
	PageInfo
}
