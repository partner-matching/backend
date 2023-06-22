package util

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport/http"
	"reflect"
	"unsafe"
)

func StructAssign(binding interface{}, value interface{}) {
	bVal := reflect.ValueOf(binding).Elem() //获取reflect.Type类型
	vVal := reflect.ValueOf(value).Elem()   //获取reflect.Type类型
	vTypeOfT := vVal.Type()
	for i := 0; i < vVal.NumField(); i++ {
		// 在要修改的结构体中查询有数据结构体中相同属性的字段，有则修改其值
		name := vTypeOfT.Field(i).Name
		if ok := bVal.FieldByName(name).IsValid(); ok {
			bVal.FieldByName(name).Set(reflect.ValueOf(vVal.Field(i).Interface()))
		}
	}
}

func GetRequestAndResponse(ctx context.Context) (*http.Request, http.ResponseWriter) {
	q := reflect.ValueOf(ctx).Elem().FieldByName("req")
	s := reflect.ValueOf(ctx).Elem().FieldByName("res")
	req := reflect.NewAt(q.Type(), unsafe.Pointer(q.UnsafeAddr())).Elem().Interface().(*http.Request)
	res := reflect.NewAt(s.Type(), unsafe.Pointer(s.UnsafeAddr())).Elem().Interface().(http.ResponseWriter)
	return req, res
}

func PrintNonZeroFieldsAndValues(s interface{}) map[string]interface{} {
	// 获取结构体的值和类型
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)
	result := make(map[string]interface{}, 0)

	// 遍历结构体的所有字段
	for i := 0; i < v.NumField(); i++ {
		// 获取字段的值和类型
		fieldValue := v.Field(i)
		fieldType := t.Field(i)

		// 检查字段是否为零值
		if !isZero(fieldValue) {
			// 打印非零值字段的名称和值
			//fmt.Printf("%s: %v\n", fieldType.Name, fieldValue.Interface())
			result[fieldType.Name] = fieldValue.Interface()
		}
	}
	return result
}

func isZero(v reflect.Value) bool {
	return v.Interface() == reflect.Zero(v.Type()).Interface()
}
