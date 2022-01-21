package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateMultiCluster
//@description: 创建MultiCluster记录
//@param: multicluster model.MultiCluster
//@return: err error

func CreateMultiCluster(multicluster model.MultiCluster) (err error) {
	err = global.GVA_DB.Create(&multicluster).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteMultiCluster
//@description: 删除MultiCluster记录
//@param: multicluster model.MultiCluster
//@return: err error

func DeleteMultiCluster(multicluster model.MultiCluster) (err error) {
	err = global.GVA_DB.Delete(&multicluster).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteMultiClusterByIds
//@description: 批量删除MultiCluster记录
//@param: ids request.IdsReq
//@return: err error

func DeleteMultiClusterByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.MultiCluster{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateMultiCluster
//@description: 更新MultiCluster记录
//@param: multicluster *model.MultiCluster
//@return: err error

func UpdateMultiCluster(multicluster model.MultiCluster) (err error) {
	err = global.GVA_DB.Save(&multicluster).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetMultiCluster
//@description: 根据id获取MultiCluster记录
//@param: id uint
//@return: err error, multicluster model.MultiCluster

func GetMultiCluster(id uint) (err error, multicluster model.MultiCluster) {
	err = global.GVA_DB.Where("id = ?", id).First(&multicluster).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetMultiClusterInfoList
//@description: 分页获取MultiCluster记录
//@param: info request.MultiClusterSearch
//@return: err error, list interface{}, total int64

func GetMultiClusterInfoList(info request.MultiClusterSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.MultiCluster{})
	var multiclusters []model.MultiCluster
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("`name` LIKE ?", "%"+info.Name+"%")
	}

	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&multiclusters).Error
	return err, multiclusters, total
}
