package utils

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"reflect"
)

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	var data = make(map[string]interface{})
	strarr := []string{"CreatedAt", "DeletedAt", "UpdatedAt", "Children", "DataAuthorityId", "Depts", "SysBaseMenus", "SysDictionaryDetails", "Status", "SysAuthoritys", "Parameters"}

	for i := 0; i < t.NumField(); i++ {
		d := v.Field(i).Interface()
		if !InArray(t.Field(i).Name, strarr) {
			switch d.(type) {
			case string:
				data[t.Field(i).Name] = "\"" + fmt.Sprintf("%v", d) + "\""
			case global.GVA_MODEL:
				d1 := d.(global.GVA_MODEL)
				data[t.Field(i).Name] = d1.ID
			case model.Meta:
				d1 := d.(model.Meta)
				data[t.Field(i).Name] = fmt.Sprintf("model.Meta{KeepAlive:%t,DefaultMenu:%t,Title:\"%s\",Icon:\"%s\"}", d1.KeepAlive, d1.DefaultMenu, d1.Title, d1.Icon)
			default:
				data[t.Field(i).Name] = d
			}
		}
	}
	return data
}
func InArray(obj interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}

	return false
}
