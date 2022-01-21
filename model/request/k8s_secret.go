package request

import "gin-vue-admin/model"

type SecretSearch struct{
    model.Secret
    PageInfo
}

type SecretSearchUser struct {
    SecretSearch
    model.User

}