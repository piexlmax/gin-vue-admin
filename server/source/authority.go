package source

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"time"

	"github.com/gookit/color"

	"gorm.io/gorm"
)

var Authority = new(authority)

type authority struct{}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: sys_authorities 表数据初始化
func (a *authority) Init() error {
	var authorities = []model.SysAuthority{
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "888", AuthorityName: i18nHash["OrdinaryUser"], ParentId: "0", DefaultRouter: "dashboard"},
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "8881", AuthorityName: i18nHash["NormalUserSubRole"], ParentId: "888", DefaultRouter: "dashboard"},
		{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "9528", AuthorityName: i18nHash["TestRole"], ParentId: "0", DefaultRouter: "dashboard"},
	}
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("authority_id IN ? ", []string{"888", "9528"}).Find(&[]model.SysAuthority{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_authorities 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&authorities).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_authorities 表初始数据成功!")
		return nil
	})
}
