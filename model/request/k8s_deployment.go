package request

import "gin-vue-admin/model"

type DeploymentSearch struct{
    model.Deployment
    PageInfo
}
type DeploymentSearchUser struct{
    DeploymentSearch
    model.User
}