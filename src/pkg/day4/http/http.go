package main

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w,"<script>alert(\"hello GO\")</script>")
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
