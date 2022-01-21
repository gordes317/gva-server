package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateStudent
//@description: 创建Student记录
//@param: student model.Student
//@return: err error

func CreateStudent(student model.Student) (err error) {
	err = global.GVA_DB.Create(&student).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteStudent
//@description: 删除Student记录
//@param: student model.Student
//@return: err error

func DeleteStudent(student model.Student) (err error) {
	err = global.GVA_DB.Delete(&student).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteStudentByIds
//@description: 批量删除Student记录
//@param: ids request.IdsReq
//@return: err error

func DeleteStudentByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Student{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateStudent
//@description: 更新Student记录
//@param: student *model.Student
//@return: err error

func UpdateStudent(student model.Student) (err error) {
	err = global.GVA_DB.Save(&student).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetStudent
//@description: 根据id获取Student记录
//@param: id uint
//@return: err error, student model.Student

func GetStudent(id uint) (err error, student model.Student) {
	err = global.GVA_DB.Where("id = ?", id).First(&student).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetStudentInfoList
//@description: 分页获取Student记录
//@param: info request.StudentSearch
//@return: err error, list interface{}, total int64

func GetStudentInfoList(info request.StudentSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Student{})
    var students []model.Student
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != "" {
        db = db.Where("`name` LIKE ?","%"+ info.Name+"%")
    }
    if info.Age != 0 {
        db = db.Where("`age` > ?",info.Age)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&students).Error
	return err, students, total
}