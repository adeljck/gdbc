package modules

import "log"

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
		log.Println("Connect Database Successful")
	} else {
		log.Fatalln("Connect Database Failed")
	}
}
func GetVersion(C Connector) {
	C.Version()
}
func GetDatabases(C Connector) {
	C.Databases()
}
