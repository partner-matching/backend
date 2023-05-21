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
