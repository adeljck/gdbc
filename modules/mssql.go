package modules

import (
	"fmt"
	"gdbc/common"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Mssql struct {
	BaseInfo common.DbInfo
	Result   struct {
		Version       string
		DataBaseInfos []MssqlDataBaseInfo
		DataBaseCount int
	}
}
type MssqlDataBaseInfo struct {
	Database   string
	TableCount int
	Tables     []string
}

var mssql *sqlx.DB

func (M *Mssql) Reverse() {

}
func (M *Mssql) Init() error {
	var err error
	M.BaseInfo = common.BaseInfo
	dsn := fmt.Sprintf("server=%s;port=%s;user id=%s;password=%s;database=%s", M.BaseInfo.Host, M.BaseInfo.Port, M.BaseInfo.UserName, M.BaseInfo.Password, "")
	mssql, err = sqlx.Open("sqlserver", dsn)
	if err != nil {
		return err
	}
	mssql.SetConnMaxLifetime(time.Minute * 3)
	mssql.SetMaxOpenConns(10)
	mssql.SetMaxIdleConns(10)
	err = mssql.Ping()
	if err != nil {
		return err
	}
	return nil
}
func (M *Mssql) Version() {
	//select Host,USER,authentication_string from mysql.user;
	err := mssql.Get(&M.Result.Version, `SELECT @@VERSION AS 'version';`)
	if err != nil {
		log.Fatalln(err)
	}
}
func (M *Mssql) Tables() {
	type t struct {
		name interface{} `mssql:"tables"`
	}
	err := mssql.Select(&M.Result.DataBaseInfos, "SELECT name  as tables FROM sys.tables;\n")
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
		err := mssql.Get(&c, query)
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

		rows, err := mssql.Query(query)
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
func (M *Mssql) Databases() {
	err := mssql.Get(&M.Result.DataBaseCount, "SELECT count(*) as count FROM sys.databases;")
	if err != nil {
		log.Fatalln(err)
	}

}
func (M *Mssql) Info() {
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
	mssql.Close()
}
func (M *Mssql) Checker() bool {
	err := M.Init()
	log.SetPrefix("[!] ")
	log.Printf("Connect To %s Server %s\n", M.BaseInfo.DbType, M.BaseInfo.Host)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
func (M Mysql) checkXpCmdShell() bool {
	t := struct {
		e string `db:"e"`
	}{}
	mssql.Get(&t, "select count(*) as e from master.dbo.sysobjects where type='x' and name='xp_cmdshell';")
	if t.e == "1" {
		return true
	}
	return false
}
