package initialize

import (
	"gin-vue-admin/global"
	"gin-vue-admin/middleware"
	"gin-vue-admin/misc/sensitive_word"
)

func WordFilter() {
	if global.GVA_CONFIG.System.UseSensitiveWordFilter {
		middleware.Library = sensitive_word.NewLibrary()
		middleware.Library.Build()
	}
}
