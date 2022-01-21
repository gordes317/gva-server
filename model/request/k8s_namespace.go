package request

import "gin-vue-admin/model"

type NamespaceSearch struct{
    model.Namespace
    PageInfo
}

type NamespaceList struct {
    model.User
    NamespaceSearch
}