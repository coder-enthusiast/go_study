package main

import (
	"encoding/json"
	"fmt"
)

type Member struct{
	Name string
	Age int
	Phone string
	Address string
}

func main()  {
	var members = make([]Member,0,10)

	for i:=0;i<10;i++ {
		var member Member
		member.Age= (i + 1)
		member.Name=fmt.Sprintf("%s%d","zzq",i)
		member.Address=fmt.Sprintf("%s%d","china",i)
		member.Phone=fmt.Sprintf("%s%d","131",i)
		members=append(members,member)
	}

	membersJson, error := json.Marshal(members)
	if error != nil {
		fmt.Println("json序列化失败",error)
	}else {
		fmt.Printf("%s\n",membersJson)
	}
	/**
	[{"name":"zzq0","age":1,"Phone":"1310","Address":"china0"},{"name":"zzq1","age":2,"Phone":"1311","Address":"china1"},{"name":"zzq2","age":3,"Phone":"1312","Address":"china2"},{"name":"zzq3","age":4,"Phone":"1313","Address":"china3"},{"name":"zzq4","age":5,"Phone":"1314","Address":"china4"},{"name":"zzq5","age":6,"Phone":"1315","Address":"china5"},{"name":"zzq6","age":7,"Phone":"1316","Address":"china6"},{"name":"zzq7","age":8,"Phone":"1317","Address":"china7"},{"name":"zzq8","age":9,"Phone":"1318","Address":"china8"},{"name":"zzq9","age":10,"Phone":"1319","Address":"china9"}]

	*/
	jsonStr := `[{"name":"zzq0","age":1,"Phone":"1310","Address":"china0"},{"name":"zzq1","age":2,"Phone":"1311","Address":"china1"},{"name":"zzq2","age":3,"Phone":"1312","Address":"china2"},{"name":"zzq3","age":4,"Phone":"1313","Address":"china3"},{"name":"zzq4","age":5,"Phone":"1314","Address":"china4"},{"name":"zzq5","age":6,"Phone":"1315","Address":"china5"},{"name":"zzq6","age":7,"Phone":"1316","Address":"china6"},{"name":"zzq7","age":8,"Phone":"1317","Address":"china7"},{"name":"zzq8","age":9,"Phone":"1318","Address":"china8"},{"name":"zzq9","age":10,"Phone":"1319","Address":"china9"}]`
	fmt.Println(jsonStr)
	ms:=[]Member{}
	err:=json.Unmarshal([]byte(jsonStr),&ms)
	if err != nil {
		fmt.Println("json反序列化失败",error)
	}else {
		for _,m := range ms {
			fmt.Println(m)
		}
	}
}
