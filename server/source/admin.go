package source

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"time"

	"github.com/gookit/color"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

var Admin = new(admin)

type admin struct{}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: sys_users 表数据初始化
func (a *admin) Init() error {
	var admins = []model.SysUser{
		{GVA_MODEL: global.GVA_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(), Username: "admin", Password: "e10adc3949ba59abbe56e057f20f883e", NickName: i18nHash["SuperAdmin"], HeaderImg: "http://qmplusimg.henrongyi.top/gva_header.jpg", AuthorityId: "888"},
		{GVA_MODEL: global.GVA_MODEL{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(), Username: "a303176530", Password: "3ec063004a6f31642261936a379fde3d", NickName: i18nHash["OtherUser"], HeaderImg: "http://qmplusimg.henrongyi.top/1572075907logo.png", AuthorityId: "9528"},
	}
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2}).Find(&[]model.SysUser{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_users 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&admins).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_users 表初始数据成功!")
		return nil
	})
}
