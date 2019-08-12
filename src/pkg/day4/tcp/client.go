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
	MsgType int
	Name string
	Msg string
}

var inputRead *bufio.Reader
var conn net.Conn
var err error




/**
 设置昵称
 */
var name string
func setName()  {
	println("请输入你的昵称进入群聊")
	input,_ := inputRead.ReadString('\n')
	inputInfo := strings.Trim(input,"\r\n")
	name=inputInfo
}

/**
心跳
 */
func isActive()  {

}
/**
请求服务器
 */
func createConn()  {
	conn,err = net.Dial("tcp","127.0.0.1:8899")
	if err != nil{
		fmt.Println("LOG:链接服务失败!",err)
		return
	}
	fmt.Println("正在链接……！")
	join()

	for{
		read:=bufio.NewReader(conn)
		var buf [128]byte
		n,err :=read.Read(buf[:])
		if err !=nil{
			fmt.Println("LOG:数据读取失败",err)
			break
		}
		str:=string(buf[:n])
		message := Msg
		json.Unmarshal([]byte(str),&message)
		switch message.MsgType {
		case 1: fmt.Println("【"+message.Name+"】加入群聊")
		case 2:fmt.Println(message.Name+"："+message.Msg)

		//心跳检测
		case 3:
			msg := Msg
			msg.Name=name
			msg.MsgType=3
			msg.Msg=message.Name
			msgjson,_:=json.Marshal(msg)
			send(conn,msgjson)
		}

	}
}

/**
发送消息
 */
func send(conn net.Conn,msg []byte)  {
	_,err := conn.Write(msg)
	if err != nil {
		fmt.Println("发送失败",err)
		return
	}
}
func join()  {
	var message = Msg
	message.MsgType=1
	message.Name=name
	msgJson, _ :=json.Marshal(message)
	_,err := conn.Write(msgJson)
	if err != nil {
		fmt.Println("LOG 链接服务器失败",err)
		return
	}
}

func main() {

	fmt.Println("=======GO 局域网群聊 V1.0===========")
	inputRead = bufio.NewReader(os.Stdin)
	for len(name)<=0{
		setName()
	}
	go createConn()

	for {
		input,_ := inputRead.ReadString('\n')
		inputInfo := strings.Trim(input,"\r\n")
		msg := Msg
		msg.Name=name
		msg.MsgType=2
		msg.Msg=inputInfo
		msgjson,_:=json.Marshal(msg)
		send(conn,msgjson)
	}

}
