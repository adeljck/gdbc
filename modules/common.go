package modules

import (
	"gdbc/common"
	"log"
)

type Connector interface {
	Checker() bool
	Reverse()
	Version()
	Tables()
	Databases()
	Info()
}

func Checker(C Connector) {
	if C.Checker() {
		log.SetPrefix("[+] ")
		log.Println("Connect Database Successful")
	} else {
		log.SetPrefix("[-] ")
		log.Fatalln("Connect Database Failed")
	}
}
func GetVersion(C Connector) {
	C.Version()
}
func GetDatabases(C Connector) {
	C.Databases()
}
func GetTables(C Connector) {
	C.Tables()
}
func PrintInfo(C Connector) {
	C.Info()
}
func CreateObject() Connector {
	switch common.BaseInfo.DbType {
	case "mysql":
		m := new(Mysql)
		return m
	default:
		log.SetPrefix("[-] ")
		log.Fatalf("UnSupport Db Type %s!!!!\n", common.BaseInfo.DbType)
		return nil
	}
}
