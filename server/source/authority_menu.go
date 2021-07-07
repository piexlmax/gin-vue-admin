package source

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"github.com/gookit/color"
)

var AuthorityMenu = new(authorityMenu)

type authorityMenu struct{}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: authority_menu 视图数据初始化
func (a *authorityMenu) Init() error {
	if global.GVA_DB.Find(&[]model.SysMenu{}).RowsAffected > 0 {
		color.Danger.Println("\n[Mysql] --> authority_menu 视图已存在!")
		return nil
	}
	if err := global.GVA_DB.Exec("CREATE VIEW authority_menu AS select sys_base_menus.id AS id,sys_base_menus.created_at AS created_at, sys_base_menus.updated_at AS updated_at, sys_base_menus.deleted_at AS deleted_at, sys_base_menus.menu_level AS menu_level,sys_base_menus.parent_id AS parent_id,sys_base_menus.path AS path,sys_base_menus.name AS name,sys_base_menus.hidden AS hidden,sys_base_menus.component AS component, sys_base_menus.title  AS title,sys_base_menus.icon AS icon,sys_base_menus.sort AS sort,sys_authority_menus.sys_authority_authority_id AS authority_id,sys_authority_menus.sys_base_menu_id AS menu_id,sys_base_menus.keep_alive AS keep_alive,sys_base_menus.default_menu AS default_menu from (sys_authority_menus join sys_base_menus on ((sys_authority_menus.sys_base_menu_id = sys_base_menus.id)))").Error; err != nil {
		return err
	}
	color.Info.Println("\n[Mysql] --> authority_menu 视图创建成功!")
	return nil
}
