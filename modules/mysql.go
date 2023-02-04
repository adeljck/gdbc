package modules

import (
	"database/sql"
	"dbconnector/common"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type Mysql struct {
	BaseInfo common.DbInfo
	Result   struct {
		Version       string
		TablesCounter []TableInfo
	}
}
type TableInfo struct {
	Database   string
	TableCount int
}

var db *sql.DB

func (M Mysql) Reverse() {

}
func (M Mysql) Init() error {
	var err error
	M.BaseInfo = common.BaseInfo
	M.BaseInfo.PrintDbInfo()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/mysql?charset=utf8mb4", M.BaseInfo.UserName, M.BaseInfo.Password, M.BaseInfo.Host, M.BaseInfo.Port)
	db, err = sql.Open(M.BaseInfo.DbType, dsn)
	if err != nil {
		return err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}
func (M Mysql) Version() {
	//select Host,USER,authentication_string from mysql.user;
	err := db.QueryRow("select version()").Scan(&M.Result.Version)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Mysql Version %s\n", M.Result.Version)
}
func (M Mysql) Tables() {

}
func (M Mysql) Databases() {
	d := ""
	rows, err := db.Query("SELECT table_schema, COUNT(1) TABLES FROM information_schema.TABLES GROUP BY table_schema;")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("database:%s\n", d)
}
func (M Mysql) Info() {

}
func (M Mysql) Checker() bool {
	err := M.Init()
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
