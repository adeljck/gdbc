package common

import "fmt"

var (
	BaseInfo    DbInfo
	DefaultPort = map[string]string{"mysql": "3306", "redis": "6379", "postgres": "5432", "oracle": "1521", "memcache": "11211", "mongodb": "27017", "mssql": "1433", "sybase": "2052", "db2": "5000"}
)

type DbInfo struct {
	DbType   string
	Host     string
	Port     string
	UserName string
	Password string
	DbUrl    string
}

func (D DbInfo) PrintDbInfo() {
	results := "Database Type: %s\nHost: %s\nPort: %s\nUser: %s\nPassword: %s\n"
	fmt.Printf(results, D.DbType, D.Host, D.Port, D.UserName, D.Password)
}
