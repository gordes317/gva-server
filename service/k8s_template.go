package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateTemplate
//@description: 创建Template记录
//@param: template model.Template
//@return: err error

func CreateTemplate(template model.Template) (err error) {
	err = global.GVA_DB.Create(&template).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTemplate
//@description: 删除Template记录
//@param: template model.Template
//@return: err error

func DeleteTemplate(template model.Template) (err error) {
	err = global.GVA_DB.Delete(&template).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTemplateByIds
//@description: 批量删除Template记录
//@param: ids request.IdsReq
//@return: err error

func DeleteTemplateByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Template{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateTemplate
//@description: 更新Template记录
//@param: template *model.Template
//@return: err error

func UpdateTemplate(template model.Template) (err error) {
	err = global.GVA_DB.Save(&template).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTemplate
//@description: 根据id获取Template记录
//@param: id uint
//@return: err error, template model.Template

func GetTemplate(id uint) (err error, template model.Template) {
	err = global.GVA_DB.Where("id = ?", id).First(&template).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTemplateInfoList
//@description: 分页获取Template记录
//@param: info request.TemplateSearch
//@return: err error, list interface{}, total int64

func GetTemplateInfoList(info request.TemplateSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Template{})
	var templates []model.Template
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("`name` LIKE ?", "%"+info.Name+"%")
	}
	if info.Type != "" {
		db = db.Where("`type` LIKE ?", "%"+info.Type+"%")
	}

	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&templates).Error
	return err, templates, total
}
