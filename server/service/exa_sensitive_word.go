package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateExaSensitiveWord
//@description: 创建ExaSensitiveWord记录
//@param: exaSensitiveWord model.ExaSensitiveWord
//@return: err error

func CreateExaSensitiveWord(exaSensitiveWord model.ExaSensitiveWord) (err error) {
	err = global.GVA_DB.Create(&exaSensitiveWord).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteExaSensitiveWord
//@description: 删除ExaSensitiveWord记录
//@param: exaSensitiveWord model.ExaSensitiveWord
//@return: err error

func DeleteExaSensitiveWord(exaSensitiveWord model.ExaSensitiveWord) (err error) {
	err = global.GVA_DB.Delete(&exaSensitiveWord).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteExaSensitiveWordByIds
//@description: 批量删除ExaSensitiveWord记录
//@param: ids request.IdsReq
//@return: err error

func DeleteExaSensitiveWordByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.ExaSensitiveWord{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateExaSensitiveWord
//@description: 更新ExaSensitiveWord记录
//@param: exaSensitiveWord *model.ExaSensitiveWord
//@return: err error

func UpdateExaSensitiveWord(exaSensitiveWord model.ExaSensitiveWord) (err error) {
	err = global.GVA_DB.Save(&exaSensitiveWord).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetExaSensitiveWord
//@description: 根据id获取ExaSensitiveWord记录
//@param: id uint
//@return: err error, exaSensitiveWord model.ExaSensitiveWord

func GetExaSensitiveWord(id uint) (err error, exaSensitiveWord model.ExaSensitiveWord) {
	err = global.GVA_DB.Where("id = ?", id).First(&exaSensitiveWord).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetExaSensitiveWordInfoList
//@description: 分页获取ExaSensitiveWord记录
//@param: info request.ExaSensitiveWordSearch
//@return: err error, list interface{}, total int64

func GetExaSensitiveWordInfoList(info request.ExaSensitiveWordSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.ExaSensitiveWord{})
    var exaSensitiveWords []model.ExaSensitiveWord
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&exaSensitiveWords).Error
	return err, exaSensitiveWords, total
}