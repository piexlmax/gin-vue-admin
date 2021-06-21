package model

import (
	"gin-vue-admin/global"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
)


func (sa *SysApi) AfterCreate(tx *gorm.DB) (err error) {
		cr:=gormadapter.CasbinRule{
			PType:       "p",
			V0:          "888",
			V1:          sa.Path,
			V2:          sa.Method,
		}
	gva := global.GVA_CONFIG.Mysql.Dbname
	if tx.Model(gormadapter.CasbinRule{}).Where(cr).First(&gormadapter.CasbinRule{}).RowsAffected > 0 {
		//global.GVA_LOG.Info("casbin already exist")
		return nil
	}
	err = tx.Exec("INSERT INTO "+gva+".casbin_rule (p_type, v0, v1, v2, v3, v4, v5) VALUES ('p', '888', ?, ?,'','','')", sa.Path, sa.Method).Error
	if err != nil {
		//fmt.Println("数据库callback sysAPI 创建：更新CASBIN失败")
		global.GVA_LOG.Error("新建API，自动授权失败", zap.Any("erro", err))
	}
	return
}
