package request

import "gin-vue-admin/model"

// Paging common input parameter structure
type PageInfo struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

// Find by id structure
type GetById struct {
	Id float64 `json:"id" form:"id"`
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

//Find by name
type NamesReq struct {
	Names []string `json:'names' form:"names"`
}
type NamesReqUser struct {
	NamesReq
	model.User
}

// Get role by id structure
type GetAuthorityId struct {
	AuthorityId string
}

type Empty struct{}
