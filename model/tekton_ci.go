package model

type Repository struct {
	PId  string `json:"PId" form:"PId"`
	ProjectName  string `json:"ProjectName" form:"ProjectName"`
	BranchName string `json:"BranchName" form:"BranchName"`
	TagName string `json:"TagName" form:"TagName"`
	GitSshUrl string `json:"gitSshUrl" form:"gitSshUrl"`
	CommitShortId  string `json:"CommitShortId" form:"CommitShortId"`
	Env  string `json:"Env" form:"Env"`
}

type Branch struct {
	PId  string `json:"PId" form:"PId"`
	BranchName string `json:"BranchName" form:"BranchName"`
	TagName string `json:"TagName" form:"TagName"`
	BranchOrTag string `json:"BranchOrTag" form:"BranchOrTag"`
}