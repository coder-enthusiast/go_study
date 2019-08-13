package main

import (
	"fmt"
	"reflect"
	"time"
)
type student struct {
	name string `sql:"NAME"`
	age  int
}
func main() {
	/*var s  student
	s.name="zzq"
	s.age=1
	data, _ := insertSql(s)
	fmt.Println(data)*/
	fmt.Println(time.Now().String())

}

/**
反射获取 字段
*/
func getPram(obj interface{}) (map[string]reflect.Value,string)  {
	ob := reflect.TypeOf(obj)
	obValue := reflect.ValueOf(obj)
	fieldKV := make(map[string]reflect.Value)
	if ob.Kind() == reflect.Struct {
		for i:=0;i<ob.NumField();i++{
			//获取字段
			f:=ob.Field(i).Name
			//获取字段value
			v := obValue.FieldByName(f)
			//获取是否有tag
			tag := ob.Field(i).Tag.Get("sql")
			if len(tag)>0 {
				f=tag
			}
			if isNotBlank(v){
				fieldKV[f] =v
			}
		}
	}
	return fieldKV,ob.Name()
}
/**
判断Value是否有值
*/
func isBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

func isNotBlank(value reflect.Value) bool {
	return !isBlank(value)
}

func insertSql(o interface{}) (string, []reflect.Value) {
	value := make([]reflect.Value,0,10)
	para,tableName := getPram(o)
	sql := "insert into "+tableName+"(%s) values (%s)"
	sqlPrefix := ""
	sqlSuffix := ""
	for k,v := range para {
		sqlPrefix+= k+","
		sqlSuffix+= "?,"
		value = append(value,v)
	}

	return fmt.Sprintf(sql,sqlPrefix[:len(sqlPrefix)-1],sqlSuffix[:len(sqlSuffix)-1]) ,value
}