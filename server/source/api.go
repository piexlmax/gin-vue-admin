package source

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"time"

	"github.com/gookit/color"

	"gorm.io/gorm"
)

var Api = new(api)

type api struct{}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: sys_apis 表数据初始化
func (a *api) Init() error {
	var apis = []model.SysApi{
		{global.GVA_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/base/login", i18nHash["UserLogin"], "base", "POST"},
		{global.GVA_MODEL{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/register", i18nHash["UserRegister"], "user", "POST"},
		{global.GVA_MODEL{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/createApi", i18nHash["CreateApi"], "api", "POST"},
		{global.GVA_MODEL{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/getApiList", i18nHash["GetApiList"], "api", "POST"},
		{global.GVA_MODEL{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/getApiById", i18nHash["GetApiDetail"], "api", "POST"},
		{global.GVA_MODEL{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/deleteApi", i18nHash["DeleteApi"], "api", "POST"},
		{global.GVA_MODEL{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/updateApi", i18nHash["UpdateApi"], "api", "POST"},
		{global.GVA_MODEL{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/getAllApis", i18nHash["GetAllApis"], "api", "POST"},
		{global.GVA_MODEL{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/createAuthority", i18nHash["CreateAuthority"], "authority", "POST"},
		{global.GVA_MODEL{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/deleteAuthority", i18nHash["DeleteAuthority"], "authority", "POST"},
		{global.GVA_MODEL{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/getAuthorityList", i18nHash["GetAuthorityList"], "authority", "POST"},
		{global.GVA_MODEL{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/getMenu", i18nHash["GetMenu"], "menu", "POST"},
		{global.GVA_MODEL{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/getMenuList", i18nHash["GetMenuList"], "menu", "POST"},
		{global.GVA_MODEL{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/addBaseMenu", i18nHash["AddBaseMenu"], "menu", "POST"},
		{global.GVA_MODEL{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/getBaseMenuTree", i18nHash["GetBaseMenuTree"], "menu", "POST"},
		{global.GVA_MODEL{ID: 16, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/addMenuAuthority", i18nHash["AddMenuAuthority"], "menu", "POST"},
		{global.GVA_MODEL{ID: 17, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/getMenuAuthority", i18nHash["GetMenuAuthority"], "menu", "POST"},
		{global.GVA_MODEL{ID: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/deleteBaseMenu", i18nHash["DeleteBaseMenu"], "menu", "POST"},
		{global.GVA_MODEL{ID: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/updateBaseMenu", i18nHash["UpdateBaseMenu"], "menu", "POST"},
		{global.GVA_MODEL{ID: 20, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/getBaseMenuById", i18nHash["GetBaseMenuById"], "menu", "POST"},
		{global.GVA_MODEL{ID: 21, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/changePassword", i18nHash["ChangePassword"], "user", "POST"},
		{global.GVA_MODEL{ID: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/getUserList", i18nHash["GetUserList"], "user", "POST"},
		{global.GVA_MODEL{ID: 24, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/setUserAuthority", i18nHash["SetUserAuthority"], "user", "POST"},
		{global.GVA_MODEL{ID: 25, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/fileUploadAndDownload/upload", i18nHash["UploadFile"], "fileUploadAndDownload", "POST"},
		{global.GVA_MODEL{ID: 26, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/fileUploadAndDownload/getFileList", i18nHash["GetFileList"], "fileUploadAndDownload", "POST"},
		{global.GVA_MODEL{ID: 27, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/casbin/updateCasbin", i18nHash["UpdateCasbin"], "casbin", "POST"},
		{global.GVA_MODEL{ID: 28, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/casbin/getPolicyPathByAuthorityId", i18nHash["GetPolicyPathByAuthorityId"], "casbin", "POST"},
		{global.GVA_MODEL{ID: 29, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/fileUploadAndDownload/deleteFile", i18nHash["DeleteFile"], "fileUploadAndDownload", "POST"},
		{global.GVA_MODEL{ID: 30, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/jwt/jsonInBlacklist", i18nHash["JsonInBlacklist"], "jwt", "POST"},
		{global.GVA_MODEL{ID: 31, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/setDataAuthority", i18nHash["SetDataAuthority"], "authority", "POST"},
		{global.GVA_MODEL{ID: 32, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/system/getSystemConfig", i18nHash["GetSystemConfig"], "system", "POST"},
		{global.GVA_MODEL{ID: 33, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/system/setSystemConfig", i18nHash["SetSystemConfig"], "system", "POST"},
		{global.GVA_MODEL{ID: 34, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/customer/customer", i18nHash["CreateCustomer"], "customer", "POST"},
		{global.GVA_MODEL{ID: 35, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/customer/customer", i18nHash["UpdateCustomer"], "customer", "PUT"},
		{global.GVA_MODEL{ID: 36, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/customer/customer", i18nHash["DeleteCustomer"], "customer", "DELETE"},
		{global.GVA_MODEL{ID: 37, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/customer/customer", i18nHash["GetCustomer"], "customer", "GET"},
		{global.GVA_MODEL{ID: 38, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/customer/customerList", i18nHash["GetCustomerList"], "customer", "GET"},
		{global.GVA_MODEL{ID: 39, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/casbin/casbinTest/:pathParam", i18nHash["RESTFULTest"], "casbin", "GET"},
		{global.GVA_MODEL{ID: 40, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/autoCode/createTemp", i18nHash["CreateTemp"], "autoCode", "POST"},
		{global.GVA_MODEL{ID: 41, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/updateAuthority", i18nHash["UpdateAuthority"], "authority", "PUT"},
		{global.GVA_MODEL{ID: 42, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/copyAuthority", i18nHash["CopyAuthority"], "authority", "POST"},
		{global.GVA_MODEL{ID: 43, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/deleteUser", i18nHash["DeleteUser"], "user", "DELETE"},
		{global.GVA_MODEL{ID: 44, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionaryDetail/createSysDictionaryDetail", i18nHash["CreateSysDictionaryDetail"], "sysDictionaryDetail", "POST"},
		{global.GVA_MODEL{ID: 45, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionaryDetail/deleteSysDictionaryDetail", i18nHash["DeleteSysDictionaryDetail"], "sysDictionaryDetail", "DELETE"},
		{global.GVA_MODEL{ID: 46, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionaryDetail/updateSysDictionaryDetail", i18nHash["UpdateSysDictionaryDetail"], "sysDictionaryDetail", "PUT"},
		{global.GVA_MODEL{ID: 47, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionaryDetail/findSysDictionaryDetail", i18nHash["FindSysDictionaryDetail"], "sysDictionaryDetail", "GET"},
		{global.GVA_MODEL{ID: 48, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionaryDetail/getSysDictionaryDetailList", i18nHash["GetSysDictionaryDetailList"], "sysDictionaryDetail", "GET"},
		{global.GVA_MODEL{ID: 49, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionary/createSysDictionary", i18nHash["CreateSysDictionary"], "sysDictionary", "POST"},
		{global.GVA_MODEL{ID: 50, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionary/deleteSysDictionary", i18nHash["DeleteSysDictionary"], "sysDictionary", "DELETE"},
		{global.GVA_MODEL{ID: 51, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionary/updateSysDictionary", i18nHash["UpdateSysDictionary"], "sysDictionary", "PUT"},
		{global.GVA_MODEL{ID: 52, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionary/findSysDictionary", i18nHash["FindSysDictionary"], "sysDictionary", "GET"},
		{global.GVA_MODEL{ID: 53, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysDictionary/getSysDictionaryList", i18nHash["GetSysDictionaryList"], "sysDictionary", "GET"},
		{global.GVA_MODEL{ID: 54, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysOperationRecord/createSysOperationRecord", i18nHash["CreateSysOperationRecord"], "sysOperationRecord", "POST"},
		{global.GVA_MODEL{ID: 55, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysOperationRecord/deleteSysOperationRecord", i18nHash["DeleteSysOperationRecord"], "sysOperationRecord", "DELETE"},
		{global.GVA_MODEL{ID: 56, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysOperationRecord/findSysOperationRecord", i18nHash["FindSysOperationRecord"], "sysOperationRecord", "GET"},
		{global.GVA_MODEL{ID: 57, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysOperationRecord/getSysOperationRecordList", i18nHash["GetSysOperationRecordList"], "sysOperationRecord", "GET"},
		{global.GVA_MODEL{ID: 58, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/autoCode/getTables", i18nHash["GetTables"], "autoCode", "GET"},
		{global.GVA_MODEL{ID: 59, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/autoCode/getDB", i18nHash["GetDB"], "autoCode", "GET"},
		{global.GVA_MODEL{ID: 60, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/autoCode/getColumn", i18nHash["GetColumn"], "autoCode", "GET"},
		{global.GVA_MODEL{ID: 61, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysOperationRecord/deleteSysOperationRecordByIds", i18nHash["DeleteSysOperationRecordByIds"], "sysOperationRecord", "DELETE"},
		{global.GVA_MODEL{ID: 62, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/simpleUploader/upload", i18nHash["SubsectionUpload"], "simpleUploader", "POST"},
		{global.GVA_MODEL{ID: 63, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/simpleUploader/checkFileMd5", i18nHash["CheckFileMd5"], "simpleUploader", "GET"},
		{global.GVA_MODEL{ID: 64, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/simpleUploader/mergeFileMd5", i18nHash["MergeFileMd5"], "simpleUploader", "GET"},
		{global.GVA_MODEL{ID: 65, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/setUserInfo", i18nHash["SetUserInfo"], "user", "PUT"},
		{global.GVA_MODEL{ID: 66, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/system/getServerInfo", i18nHash["GetServerInfo"], "system", "POST"},
		{global.GVA_MODEL{ID: 67, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/email/emailTest", i18nHash["EmailTest"], "email", "POST"},
		{global.GVA_MODEL{ID: 68, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/createWorkflowProcess", i18nHash["CreateWorkflowProcess"], "workflowProcess", "POST"},
		{global.GVA_MODEL{ID: 69, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/deleteWorkflowProcess", i18nHash["DeleteWorkflowProcess"], "workflowProcess", "DELETE"},
		{global.GVA_MODEL{ID: 70, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/deleteWorkflowProcessByIds", i18nHash["DeleteWorkflowProcessByIds"], "workflowProcess", "DELETE"},
		{global.GVA_MODEL{ID: 71, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/updateWorkflowProcess", i18nHash["UpdateWorkflowProcess"], "workflowProcess", "PUT"},
		{global.GVA_MODEL{ID: 72, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/findWorkflowProcess", i18nHash["FindWorkflowProcess"], "workflowProcess", "GET"},
		{global.GVA_MODEL{ID: 73, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/getWorkflowProcessList", i18nHash["GetWorkflowProcessList"], "workflowProcess", "GET"},
		{global.GVA_MODEL{ID: 74, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/findWorkflowStep", i18nHash["FindWorkflowStep"], "workflowProcess", "GET"},
		{global.GVA_MODEL{ID: 75, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/startWorkflow", i18nHash["StartWorkflow"], "workflowProcess", "POST"},
		{global.GVA_MODEL{ID: 76, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/getMyStated", i18nHash["GetMyStated"], "workflowProcess", "GET"},
		{global.GVA_MODEL{ID: 77, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/getMyNeed", i18nHash["GetMyNeed"], "workflowProcess", "GET"},
		{global.GVA_MODEL{ID: 78, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/getWorkflowMoveByID", i18nHash["GetWorkflowMoveByID"], "workflowProcess", "GET"},
		{global.GVA_MODEL{ID: 79, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/workflowProcess/completeWorkflowMove", i18nHash["CompleteWorkflowMove"], "workflowProcess", "POST"},
		{global.GVA_MODEL{ID: 80, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/autoCode/preview", i18nHash["Preview"], "autoCode", "POST"},
		{global.GVA_MODEL{ID: 81, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/excel/importExcel", i18nHash["ImportExcel"], "excel", "POST"},
		{global.GVA_MODEL{ID: 82, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/excel/loadExcel", i18nHash["LoadExcel"], "excel", "GET"},
		{global.GVA_MODEL{ID: 83, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/excel/exportExcel", i18nHash["ExportExcel"], "excel", "POST"},
		{global.GVA_MODEL{ID: 84, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/excel/downloadTemplate", i18nHash["DownloadTemplate"], "excel", "GET"},
	}
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 67}).Find(&[]model.SysApi{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_apis 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&apis).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_apis 表初始数据成功!")
		return nil
	})
}
