package utils

import (
	"fmt"
	"github.com/go-ini/ini"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	AccessKey   string
	SecretKey   string
	Bucket      string
	QiniuServer string
)

func init() {
	file, err := ini.Load("C:\\Users\\Administrator\\Desktop\\GoBlog\\config\\config.ini")
	//file, err := ini.Load("D:\\GoBlog\\config\\config_home.ini")

	if err != nil {
		fmt.Println("配置文件错误,请检查文件路径", err)
	}
	LoadServer(file)
	LoadData(file)
	// TODO 视频里面没有load这个创建的函数 我忘记打了
	LoadQiniu(file)

}

func LoadServer(file *ini.File) {
	// 这个地方就是再ini文件中写的分区
	// 其类似key value的格式
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("liushiyishishabi123")
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("8888")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("123456")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")

}

func LoadQiniu(file *ini.File) {
	AccessKey = file.Section("qiniu").Key("AccessKey").MustString("GgZMU9eK9F2QpgH1qGJEAjeI5dXywCJDdGYBT14e")
	SecretKey = file.Section("qiniu").Key("SecretKey").MustString("mKsjGKePRYbJKlts_B8Ln4ThTcsu-v-TdaiKGb39")
	Bucket = file.Section("qiniu").Key("Bucket").MustString("wangzijianblog")
	QiniuServer = file.Section("qiniu").Key("QiniuServer").MustString("http://s61pj5xc1.hn-bkt.clouddn.com/")
}
