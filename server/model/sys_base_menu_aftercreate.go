package model

import (
	"gin-vue-admin/global"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SysAuthorityMenus struct {
	SysAuthorityAuthorityId string
	SysBaseMenuId           uint
}

func (sy *SysBaseMenu) AfterCreate(tx *gorm.DB) (err error) {
	var sam =SysAuthorityMenus{
		SysAuthorityAuthorityId: "888",
		SysBaseMenuId:           sy.ID,
	}
	if tx.Table("sys_authority_menus").Where("  sys_authority_authority_id = '888' and sys_base_menu_id= ? ", sy.ID).Scan(&SysAuthorityMenus{}).RowsAffected > 0 {
		return nil
	}
	err = tx.Table("sys_authority_menus").Create(&sam).Error
	if err != nil {
		global.GVA_LOG.Error("新建menu，自动授权失败", zap.Any("erro", err))
	} else {
		global.GVA_LOG.Info("新建menu，自动授权成功")
	}
	return nil
}
