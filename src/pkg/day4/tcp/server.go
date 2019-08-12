package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
)


var Msg struct{
	MsgType int //1加入群聊 2正常消息 3心跳检测
	Name string
	Msg string
}

/**
socket
 */
var isSatrt=false

var conns =make(map[string]net.Conn)
func main() {
	fmt.Println("==============GO 局域网聊天系统 服务器端===============")

	menu()
	inputRead := bufio.NewReader(os.Stdin)
	for {
		input,_ := inputRead.ReadString('\n')
		inputInfo := strings.Trim(input,"\r\n")
		switch inputInfo {
		case "1":go start()
			     isSatrt=true
		case "2":os.Exit(0)
		case "3":fmt.Println("当前用户个数：",len(conns))
		case "4":fmt.Println("========用户列表==========")
			for k,_ := range conns {
				fmt.Println(k)
			}


		case "0":menu()
		}

	}




}

var listen  net.Listener
func start()  {
	listen,err :=net.Listen("tcp","127.0.0.1:8899")
	if err!=nil{
		fmt.Println("LOG:服务器创建失败",err)
		return
	}
	fmt.Println("LOG:服务器启动完毕！")
	for   {
		//监听链接
		conn,err :=listen.Accept()
		if err!=nil{
			fmt.Println("LOG:建立链接失败！",err)
			continue
		}
		// 线程处理链接
		go process(conn)
	}
}


func process(conn net.Conn)  {
	//关闭链接
	defer conn.Close()
	for{
		read:=bufio.NewReader(conn)
		var buf [128]byte
		n,err :=read.Read(buf[:])
		if err !=nil{
			fmt.Println("LOG:数据读取失败",err)
			break
		}
		str:=string(buf[:n])
		//fmt.Println("客户端发来的内容：",str)
		message := Msg
		json.Unmarshal([]byte(str),&message)
		switch message.MsgType {
		case 1:
			conns[message.Name]=conn
			for _,v := range conns {
				v.Write([]byte(str))

		}

		case 2:
			for k,v := range conns {
				if k!=message.Name {
					v.Write([]byte(str))
				}
			}

		}
	}
}

func menu()  {
	if !isSatrt {
		fmt.Println("1.启动服务器")
	}
	fmt.Println("2.退出")
	fmt.Println("3.当前用户个数")
	fmt.Println("4.当前用户列表")
	fmt.Println("0.菜单")


}

