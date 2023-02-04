package modules

import (
	"fmt"
	"gdbc/common"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Mysql struct {
	BaseInfo common.DbInfo
	Result   struct {
		Version       string `db:"version"`
		DataBaseInfos []DataBaseInfo
		DataBaseCount int `db:"count"`
	}
}
type DataBaseInfo struct {
	Database   string `db:"table_schema"`
	TableCount int    `db:"TABLES"`
	Tables     []string
}

var db *sqlx.DB

// Reverse TODO
func (M *Mysql) Reverse() {

}
func (M *Mysql) Init() error {
	var err error
	M.BaseInfo = common.BaseInfo
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/mysql?charset=utf8mb4", M.BaseInfo.UserName, M.BaseInfo.Password, M.BaseInfo.Host, M.BaseInfo.Port)
	db, err = sqlx.Open(M.BaseInfo.DbType, dsn)
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
func (M *Mysql) Version() {
	//select Host,USER,authentication_string from mysql.user;
	err := db.Get(&M.Result, "select version() as version")
	if err != nil {
		log.Fatalln(err)
	}
}
func (M *Mysql) Tables() {
	type t struct {
		name interface{} `db:"tables"`
	}
	err := db.Select(&M.Result.DataBaseInfos, "SELECT table_schema, COUNT(1) TABLES FROM information_schema.TABLES GROUP BY table_schema;")
	if err != nil {
		log.Fatalln(err)
	}
	for i, v := range M.Result.DataBaseInfos {
		count := 20
		a := ""
		query := fmt.Sprintf("select count(*) as count from (select TABLE_NAME as tables from information_schema.tables where table_schema = '%s' ) as x;", v.Database)
		c := struct {
			Count int `db:"count"`
		}{}
		err := db.Get(&c, query)
		if err != nil {
			log.Fatalln(err)
		}
		if c.Count > count {
			fmt.Printf("[!] Database %s more than 20 tables,still get all of them?(y/n)", v.Database)
			fmt.Scanf("%s\n", &a)
		}
		if strings.ToLower(a) == "y" || strings.ToLower(a) == "" {
			count = c.Count
		}
		query = fmt.Sprintf("select TABLE_NAME as tables from information_schema.tables where table_schema = '%s';", v.Database)

		rows, err := db.Query(query)
		if err != nil {
			log.Fatalln(err)
		}
		for rows.Next() {
			count = count - 1
			var table string
			err = rows.Scan(&table)
			if err != nil {
				log.Fatalln(err)
			}
			M.Result.DataBaseInfos[i].Tables = append(M.Result.DataBaseInfos[i].Tables, table)
			if count < 0 {
				break
			}
		}
		err = rows.Err()
	}
}
func (M *Mysql) Databases() {
	err := db.Get(&M.Result, "select count(*) as count from (SELECT table_schema, COUNT(1) TABLES FROM information_schema.TABLES GROUP BY table_schema) as x;")
	if err != nil {
		log.Fatalln(err)
	}
}
func (M *Mysql) Info() {
	fmt.Println("********************************************************************************************")
	results := "Database Type: %s\nDatabase Version: %s\nHost: %s\nPort: %s\nUser: %s\nPassword: %s\n"
	fmt.Printf(results, M.BaseInfo.DbType, M.Result.Version, M.BaseInfo.Host, M.BaseInfo.Port, M.BaseInfo.UserName, M.BaseInfo.Password)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"database_name", "tables_count", "tables"})
	for _, v := range M.Result.DataBaseInfos {
		row := []string{
			v.Database, strconv.Itoa(v.TableCount), strings.Join(v.Tables, "\n"),
		}
		table.Append(row)
	}
	table.SetRowLine(true)
	table.SetCenterSeparator("*")
	table.SetColumnSeparator("|")
	table.SetRowSeparator("-")
	table.Render()
}
func (M *Mysql) Checker() bool {
	err := M.Init()
	log.SetPrefix("[!] ")
	log.Printf("Connect To %s Server %s\n", M.BaseInfo.DbType, M.BaseInfo.Host)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
