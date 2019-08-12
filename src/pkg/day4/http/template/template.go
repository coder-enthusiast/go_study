package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request)  {
	tmpl,err:=template.ParseFiles("hello.html")
	if err != nil {
		fmt.Println("err",err)
		return
	}
	tmpl.ExecuteTemplate(w,"word","HELLO GO!")

}

func main() {
	//路由
	http.HandleFunc("/",sayHello)
	err :=http.ListenAndServe(":9999",nil)
	if   err != nil {
		fmt.Println("建立http服务器失败",err)
		return
	}
}
