package common

import (
	"flag"
	"log"
	"strings"
)

func InitArgs() {
	flag.StringVar(&BaseInfo.DbUrl, "U", "", "database url example:mysql://root:123456@127.0.0.1:3306/test")
	flag.StringVar(&BaseInfo.DbType, "d", "", "database type example:mysql,mssql,oracle")
	flag.StringVar(&BaseInfo.Host, "h", "", "database host example:127.0.0.1")
	flag.StringVar(&BaseInfo.Port, "P", "", "database port example:3306(optional,with default value.)")
	flag.StringVar(&BaseInfo.UserName, "u", "", "database user example:root")
	flag.StringVar(&BaseInfo.Password, "p", "", "database password example:123456")
	flag.Parse()
	checkArgs()
}
func checkArgs() {
	if BaseInfo.DbUrl != "" {
		splitDatabaseUrl(BaseInfo.DbUrl)
		return
	} else {
		if BaseInfo.DbType == "" || BaseInfo.Host == "" || BaseInfo.UserName == "" {
			log.Fatalln("No Enough Db Info......")
		}
		BaseInfo.DbType = strings.ToLower(BaseInfo.DbType)
		BaseInfo.Port = DefaultPort[BaseInfo.DbType]
	}
}

func splitDatabaseUrl(DbUrl string) {
	first := strings.Split(DbUrl, "://")
	BaseInfo.DbType = strings.ToLower(first[0])
	if strings.Count(first[1], "@") != 1 {

	} else {
		second := strings.Split(first[1], "@")
		credentials := strings.Split(second[0], ":")
		BaseInfo.UserName = credentials[0]
		BaseInfo.Password = credentials[1]
		hostInfo := strings.Split(second[1], "/")[0]
		host := strings.Split(hostInfo, ":")
		BaseInfo.Host = host[0]
		if len(host) > 1 {
			BaseInfo.Port = host[1]
		} else {
			BaseInfo.Port = DefaultPort[BaseInfo.DbType]
		}
	}
}
