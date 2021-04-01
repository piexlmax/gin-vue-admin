package request

import "gin-vue-admin/model"

type ExaSensitiveWordSearch struct{
    model.ExaSensitiveWord
    PageInfo
}