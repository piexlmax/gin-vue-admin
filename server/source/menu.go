package source

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"time"

	"github.com/gookit/color"

	"gorm.io/gorm"
)

var BaseMenu = new(menu)

type menu struct{}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: sys_base_menus 表数据初始化
func (m *menu) Init() error {
	var menus = []model.SysBaseMenu{
		{GVA_MODEL: global.GVA_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "dashboard", Name: "dashboard", Hidden: false, Component: "view/dashboard/index.vue", Sort: 1, Meta: model.Meta{Title: i18nHash["InstrumentPanel"], Icon: "setting"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "about", Name: "about", Component: "view/about/index.vue", Sort: 7, Meta: model.Meta{Title: i18nHash["AboutUs"], Icon: "info"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "admin", Name: "superAdmin", Component: "view/superAdmin/index.vue", Sort: 3, Meta: model.Meta{Title: i18nHash["SuperAdmin"], Icon: "user-solid"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "authority", Name: "authority", Component: "view/superAdmin/authority/authority.vue", Sort: 1, Meta: model.Meta{Title: i18nHash["RoleManagement"], Icon: "s-custom"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "menu", Name: "menu", Component: "view/superAdmin/menu/menu.vue", Sort: 2, Meta: model.Meta{Title: i18nHash["MenuManagement"], Icon: "s-order", KeepAlive: true}},
		{GVA_MODEL: global.GVA_MODEL{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "api", Name: "api", Component: "view/superAdmin/api/api.vue", Sort: 3, Meta: model.Meta{Title: i18nHash["APIManagement"], Icon: "s-platform", KeepAlive: true}},
		{GVA_MODEL: global.GVA_MODEL{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "user", Name: "user", Component: "view/superAdmin/user/user.vue", Sort: 4, Meta: model.Meta{Title: i18nHash["UserManagement"], Icon: "coordinate"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: true, ParentId: "0", Path: "person", Name: "person", Component: "view/person/person.vue", Sort: 4, Meta: model.Meta{Title: i18nHash["PersonalInformation"], Icon: "message-solid"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "example", Name: "example", Component: "view/example/index.vue", Sort: 6, Meta: model.Meta{Title: i18nHash["SampleFile"], Icon: "s-management"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "excel", Name: "excel", Component: "view/example/excel/excel.vue", Sort: 4, Meta: model.Meta{Title: i18nHash["ExcelImportExport"], Icon: "s-marketing"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "upload", Name: "upload", Component: "view/example/upload/upload.vue", Sort: 5, Meta: model.Meta{Title: i18nHash["MediaLibrary"], Icon: "upload"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "breakpoint", Name: "breakpoint", Component: "view/example/breakpoint/breakpoint.vue", Sort: 6, Meta: model.Meta{Title: i18nHash["BreakpointContinuation"], Icon: "upload"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "customer", Name: "customer", Component: "view/example/customer/customer.vue", Sort: 7, Meta: model.Meta{Title: i18nHash["CustomerList"], Icon: "s-custom"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "systemTools", Name: "systemTools", Component: "view/systemTools/index.vue", Sort: 5, Meta: model.Meta{Title: i18nHash["SystemTool"], Icon: "s-cooperation"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "14", Path: "autoCode", Name: "autoCode", Component: "view/systemTools/autoCode/index.vue", Sort: 1, Meta: model.Meta{Title: i18nHash["CodeGenerator"], Icon: "cpu", KeepAlive: true}},
		{GVA_MODEL: global.GVA_MODEL{ID: 16, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "14", Path: "formCreate", Name: "formCreate", Component: "view/systemTools/formCreate/index.vue", Sort: 2, Meta: model.Meta{Title: i18nHash["FormGenerator"], Icon: "magic-stick", KeepAlive: true}},
		{GVA_MODEL: global.GVA_MODEL{ID: 17, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "14", Path: "system", Name: "system", Component: "view/systemTools/system/system.vue", Sort: 3, Meta: model.Meta{Title: i18nHash["SystemConfiguration"], Icon: "s-operation"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "dictionary", Name: "dictionary", Component: "view/superAdmin/dictionary/sysDictionary.vue", Sort: 5, Meta: model.Meta{Title: i18nHash["DictionaryManagement"], Icon: "notebook-2"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: true, ParentId: "3", Path: "dictionaryDetail/:id", Name: "dictionaryDetail", Component: "view/superAdmin/dictionary/sysDictionaryDetail.vue", Sort: 1, Meta: model.Meta{Title: i18nHash["DictionaryDetails"], Icon: "s-order"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 20, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "operation", Name: "operation", Component: "view/superAdmin/operation/sysOperationRecord.vue", Sort: 6, Meta: model.Meta{Title: i18nHash["OperationHistory"], Icon: "time"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 21, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "simpleUploader", Name: "simpleUploader", Component: "view/example/simpleUploader/simpleUploader", Sort: 6, Meta: model.Meta{Title: i18nHash["BreakpointContinuationPlug"], Icon: "upload"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 22, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "https://www.gin-vue-admin.com", Name: "https://www.gin-vue-admin.com", Hidden: false, Component: "/", Sort: 0, Meta: model.Meta{Title: i18nHash["OfficialWebsite"], Icon: "s-home"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "state", Name: "state", Hidden: false, Component: "view/system/state.vue", Sort: 6, Meta: model.Meta{Title: i18nHash["ServerStatus"], Icon: "cloudy"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 24, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "workflow", Name: "workflow", Hidden: false, Component: "view/workflow/index.vue", Sort: 5, Meta: model.Meta{Title: i18nHash["WorkflowFunction"], Icon: "phone"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 25, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "24", Path: "workflowCreate", Name: "workflowCreate", Hidden: false, Component: "view/workflow/workflowCreate/workflowCreate.vue", Sort: 0, Meta: model.Meta{Title: i18nHash["WorkflowDrawing"], Icon: "circle-plus"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 26, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "24", Path: "workflowProcess", Name: "workflowProcess", Hidden: false, Component: "view/workflow/workflowProcess/workflowProcess.vue", Sort: 0, Meta: model.Meta{Title: i18nHash["WorkflowList"], Icon: "s-cooperation"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 27, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "24", Path: "workflowUse", Name: "workflowUse", Hidden: true, Component: "view/workflow/workflowUse/workflowUse.vue", Sort: 0, Meta: model.Meta{Title: i18nHash["UseWorkflow"], Icon: "video-play"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 28, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "24", Path: "started", Name: "started", Hidden: false, Component: "view/workflow/userList/started.vue", Sort: 0, Meta: model.Meta{Title: i18nHash["IInitiatedIt"], Icon: "s-order"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 29, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "24", Path: "need", Name: "need", Hidden: false, Component: "view/workflow/userList/need.vue", Sort: 0, Meta: model.Meta{Title: i18nHash["MyToDoList"], Icon: "s-platform"}},
	}
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 29}).Find(&[]model.SysBaseMenu{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_base_menus 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&menus).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_base_menus 表初始数据成功!")
		return nil
	})
}
