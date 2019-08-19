package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/guregu/null.v3"
	"gopkg.in/guregu/null.v3/zero"
	"reflect"
	"strconv"
)

type Dto struct {
	Id null.Int
	Name null.String
	Create null.Time
	IsOnLine null.Bool
	Weight null.Float
	Test zero.Bool
}
const(
	TIME  = "time.Time"
)

func main() {
	var dto Dto
	dto.Id=null.NewInt(1,true)


	dto.Name=null.NewString("朱志强",false)

	dto.Test =zero.NewBool(true,false)
	fmt.Println(dto)

	jsonStr, _ := json.Marshal(dto)
	fmt.Println(string(jsonStr))
	fmt.Println(reflect.TypeOf(null.Int{}).String())
	fmt.Println(reflect.TypeOf(null.String{}).String())
	fmt.Println(reflect.TypeOf(null.Time{}).String())
	fmt.Println(reflect.TypeOf(null.Float{}).String())
	fmt.Println(reflect.TypeOf(null.Bool{}).String())



	fmt.Println(reflect.TypeOf(zero.Int{}).String())
	fmt.Println(reflect.TypeOf(zero.String{}).String())
	fmt.Println(reflect.TypeOf(zero.Time{}).String())
	fmt.Println(reflect.TypeOf(zero.Float{}).String())
	fmt.Println(reflect.TypeOf(zero.Bool{}).String())


	fmt.Println(strconv.ParseFloat("",64))


}
