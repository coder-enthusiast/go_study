package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"reflect"
)

var db *sql.DB


type A struct{
	id int
	name string
	address string
	age int64
}

func initDb() error  {
	dns := "root:123456@tcp(127.0.0.1:3306)/go_test"
	var err error
	db,err = sql.Open("mysql",dns)
	if err!=nil {
		logrus.WithFields(logrus.Fields{}).Error("数据库链接失败")
		return err
	}
	err = db.Ping()
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("PING ERR")
		return err
	}
	//设置最大连接数 0不限制
	db.SetMaxOpenConns(0)
	//设置最大闲置连接数
	db.SetMaxIdleConns(10)

	return err

}

func main() {
	err := initDb()
	if err!=nil{
		logrus.WithFields(logrus.Fields{}).Error(err)
		return
	}
	/*a:= getById(1)
	fmt.Println(a)
	all := getAll()
	fmt.Println(all)
	//fmt.Println(save())
	var a A
	//a.name="朱志强"
	a.id=10
	//a.age=18
	//a.address="qweerr"
	sql,pa := getDeleteSql(a,"id")
	fmt.Println(sql,pa)*/
	save()
}

func save() int64 {
	var a A
	a.name="朱志强"
	//sql,para := insertSql(a)
//insert into "+tableName+"(%s) values (%s)
	stmt, err := db.Prepare("insert into a(name,`create`) values (?,?)")
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
	}
	defer stmt.Close()

	result,err := stmt.Exec("朱志强","2019-08-13 17:28:55")
	if err !=nil {
		logrus.WithFields(logrus.Fields{}).Error(err)
		return 0
	}
	re,_ := result.RowsAffected()
	return re
}

func getDeleteSql(o interface{},sqlwhere ...string) (string, []interface{})  {
	fieldMap := make(map[string]int)
	wherePara := make([]interface{},0,5)

	value := make([]interface{},0,10)
	para,tableName := getPram(o)
	sql := "delete from "+tableName+" %s "

	sqlSuffix := "where 1=1 "
	for _,s := range sqlwhere {
		sqlSuffix+= "and "+s+ " = ?  "
		if(fieldMap[s]<=0){
			fieldMap[s]=1
		}

	}
	for k,v := range para {
		if fieldMap[k] <=0  {

		}else {
			wherePara=append(wherePara,v)
		}

	}
	value=wherePara

	return fmt.Sprintf(sql,sqlSuffix[:len(sqlSuffix)]) ,value
}
func getUpdate(o interface{},where ...string) (string, []interface{})  {
	fieldMap := make(map[string]int)
	wherePara := make([]interface{},0,5)

	value := make([]interface{},0,10)
	para,tableName := getPram(o)
	sql := "update "+tableName+" set %s  %s"
	sqlPrefix := ""
	sqlSuffix := "where 1=1 "
	for _,s := range where {
		sqlSuffix+= "and "+s+ " = ?  "
		if(fieldMap[s]<=0){
			fieldMap[s]=1
		}
	}


	for k,v := range para {
		if fieldMap[k] <=0  {
			sqlPrefix+= k+" = ? ,"
			value = append(value,v)
		}else {
			wherePara=append(wherePara,v)
		}

	}
	value=append(value,wherePara...)

	return fmt.Sprintf(sql,sqlPrefix[:len(sqlPrefix)-1],sqlSuffix[:len(sqlSuffix)]) ,value
}


func getAll() []A {
	list := make([]A,0,10)

	row,err := db.Query("select * from a ")
	if err != nil{
		logrus.WithFields(logrus.Fields{}).Error("查询失败")
	}
	for row.Next() {

		var t A
		//一定要Scan 不然链接不释放
		err := row.Scan(&t.id,&t.name)

		if err != nil{
			logrus.WithFields(logrus.Fields{}).Error("封装失败")
		}
		list = append(list,t)
	}
	return list
}

func getById(id int) A {
	row := db.QueryRow("select * from a where id=?",id)

	var t A
	//一定要Scan 不然链接不释放
	err := row.Scan(&t.id,&t.name)

	if err != nil{
		logrus.WithFields(logrus.Fields{}).Error("封装失败")
	}
	return t
}



/**
反射获取 字段
*/
func getPram(obj interface{}) (map[string]interface{},string)  {
	ob := reflect.TypeOf(obj)
	obValue := reflect.ValueOf(obj)
	fieldKV := make(map[string]interface{})
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
				fieldKV[f] =conver(v)
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

func insertSql(o interface{}) (string, []interface{}) {
	value := make([]interface{},0,10)
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

func conver(value reflect.Value) interface{}  {
	switch value.Kind() {
	case reflect.String:
		return value.String()
	case reflect.Bool:
		return value.Int()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint()
	case reflect.Float32, reflect.Float64:
		return value.Float()
	default:
		panic("只支持基本类型")
	}
}