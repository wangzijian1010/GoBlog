package utils

import (
	"fmt"
	"github.com/go-ini/ini"
)

var (
	AppMode  string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func init() {
	file, err := ini.Load("C:\\Users\\Administrator\\Desktop\\GoBlog\\config\\config.ini")
	//file, err := ini.Load("D:\\GoBlog\\config\\config_home.ini")

	if err != nil {
		fmt.Println("配置文件错误,请检查文件路径", err)
	}
	LoadServer(file)
	LoadData(file)

}

func LoadServer(file *ini.File) {
	// 这个地方就是再ini文件中写的分区
	// 其类似key value的格式
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("8888")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("123456")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")

}
