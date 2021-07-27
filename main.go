package main

import (
	"./conf"
	"fmt"
)

func main()  {

	fmt.Println("test conf")

	testINI()
	testJSON()
	testXML()
	testYAML()

	testJsonStruct()
	testXmlStruct()

	testCnf()
	testRedisConf()
}
//  redis.conf 文件解析
func testRedisConf()  {
	fmt.Println("-----------------------------------------")
	fmt.Println("redis.conf")

	i := conf.New("./config_file/redis/redis.conf")
	fmt.Println(i)
	fmt.Println(i.Get("appendfilename"))
}
// mysql.cnf 文件解析
func testCnf()  {
	fmt.Println("-----------------------------------------")
	fmt.Println("mysql.cnf")

	i := conf.New("./config_file/mysql/mysql_simple.cnf")
	fmt.Println(i)
	fmt.Println(i.Get("mysqld"))
}
// ini 格式配置文件解析
func testINI()  {

	fmt.Println("-----------------------------------------")
	fmt.Println("ini")

	i := conf.New("./example_file/conf.ini")
	fmt.Println(i)
	fmt.Println(i.Get("server", "host"))
}

// JSON 文件解析
func testJSON()  {

	fmt.Println("-----------------------------------------")
	fmt.Println("json")

	j := conf.New("./example_file/conf.json")
	fmt.Println(j)
	fmt.Println(j.Get("server", "host"))
}
// 绑定到结构体
func testJsonStruct()  {

	fmt.Println("-----------------------------------------")
	fmt.Println("json struct")

	type Server struct {
		Host string 	`json:"host"`
		Port int 		`json:"port"`
		Dbs  []int 		`json:"dbs"`
	}
	
	type Config struct {
		Debug    bool   `json:"debug"`
		Server   Server `json:"server"`
	}

	config := &Config{}
	conf.NewJson("./example_file/conf.json", config)

	fmt.Println(config)
}
// xml  文件解析
func testXML()  {

	fmt.Println("-----------------------------------------")
	fmt.Println("xml")

	x := conf.New("./example_file/conf.xml")
	fmt.Println(x)
	fmt.Println(x.Get("config", "server", "host"))

}
// 绑定到结构体
func testXmlStruct()  {
	fmt.Println("-----------------------------------------")
	fmt.Println("xml struct")

	type Server struct {
		Host string 	`xml:"host"`
		Port int 		`xml:"port"`
		Dbs  string 	`xml:"dbs"`
	}

	type Config struct {
		Debug    bool   `xml:"debug"`
		Server   Server `xml:"server"`
	}

	config := &Config{}
	conf.NewXml("./example_file/conf.xml", config)

	fmt.Println(config)
}

// YAML 文件解析
func testYAML()  {

	fmt.Println("-----------------------------------------")
	fmt.Println("yaml")

	y := conf.New("./example_file/conf.yaml")
	fmt.Println(y)
	fmt.Println(y.Get("server", "master", "host"))
}