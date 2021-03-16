package source

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"time"

	"github.com/gookit/color"

	"gorm.io/gorm"
)

var Dictionary = new(dictionary)

type dictionary struct{}

var status = new(bool)

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: sys_dictionaries 表数据初始化
func (d *dictionary) Init() error {
	*status = true
	var dictionaries = []model.SysDictionary{
		{GVA_MODEL: global.GVA_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: i18nHash["Sex"], Type: "sex", Status: status, Desc: i18nHash["SexDictionary"]},
		{GVA_MODEL: global.GVA_MODEL{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: i18nHash["DBTypeInt"], Type: "int", Status: status, Desc: i18nHash["DBTypeInt"]},
		{GVA_MODEL: global.GVA_MODEL{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: i18nHash["DBTypeDateTime"], Type: "time.Time", Status: status, Desc: i18nHash["DBTypeDateTime"]},
		{GVA_MODEL: global.GVA_MODEL{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: i18nHash["DBTypeFloat"], Type: "float64", Status: status, Desc: i18nHash["DBTypeFloat"]},
		{GVA_MODEL: global.GVA_MODEL{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: i18nHash["DBTypeString"], Type: "string", Status: status, Desc: i18nHash["DBTypeString"]},
		{GVA_MODEL: global.GVA_MODEL{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: i18nHash["DBTypeBool"], Type: "bool", Status: status, Desc: i18nHash["DBTypeBool"]},
	}
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 6}).Find(&[]model.SysDictionary{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_dictionaries 表初始数据已存在!")
			return nil
		}
		if err := tx.Create(&dictionaries).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_dictionaries 表初始数据成功!")
		return nil
	})
}
