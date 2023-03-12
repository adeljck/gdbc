package modules

import (
	"context"
	"errors"
	"fmt"
	"gdbc/common"
	"github.com/olekukonko/tablewriter"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Redis struct {
	BaseInfo common.DbInfo
	Result   struct {
		Version       string `db:"version"`
		DataBaseInfos []RedisDatabaseInfo
		DataBaseCount int `db:"count"`
	}
	rs *redis.Client
}
type RedisDatabaseInfo struct {
	Database  int
	KeysCount int
	Keys      map[string]string
}

func (R *Redis) init() error {
	R.BaseInfo = common.BaseInfo
	dsn := fmt.Sprintf("%s:%s", R.BaseInfo.Host, R.BaseInfo.Port)
	R.rs = redis.NewClient(
		&redis.Options{Addr: dsn, Password: R.BaseInfo.Password, DB: 0, DialTimeout: 5 * time.Second},
	)
	pong, err := R.rs.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	if pong == "PONG" {
		return nil
	} else {
		return errors.New(fmt.Sprintf("Connect to Redis Server %s Failed!!!\n", R.BaseInfo.Host))
	}
}
func (R *Redis) Checker() bool {
	err := R.init()
	log.SetPrefix("[!] ")
	log.Printf("Connect To Redis Server %s....\n", R.BaseInfo.Host)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
func (D RedisDatabaseInfo) combineKeyVal() string {
	results := make([]string, 0)
	for k, v := range D.Keys {
		results = append(results, k+":"+v)
	}
	return strings.Join(results, "\n")
}
func (R *Redis) Info() {
	fmt.Println("********************************************************************************************")
	results := "Database Type: %s\nDatabase Version: %s\nHost: %s\nPort: %s\nUser: %s\nPassword: %s\n"
	fmt.Printf(results, R.BaseInfo.DbType, R.Result.Version, R.BaseInfo.Host, R.BaseInfo.Port, R.BaseInfo.UserName, R.BaseInfo.Password)
	if len(R.Result.DataBaseInfos) == 0 {
		fmt.Println("No Keys Found.")
		return
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"database_name", "keys_count", "key:value"})
	for _, v := range R.Result.DataBaseInfos {
		row := []string{
			strconv.Itoa(v.Database), strconv.Itoa(v.KeysCount), v.combineKeyVal(),
		}
		table.Append(row)
	}
	table.SetRowLine(true)
	table.SetCenterSeparator("*")
	table.SetColumnSeparator("|")
	table.SetRowSeparator("-")
	table.Render()
	R.rs.Close()
}
func (R *Redis) Version() {
	info, err := R.rs.Info(context.Background(), "server").Result()
	if err != nil {
		log.Fatalln(err)
	}
	version := ""
	for _, line := range strings.Split(info, "\r\n") {
		if strings.HasPrefix(line, "redis_version:") {
			version = strings.TrimPrefix(line, "redis_version:")
			break
		}
	}
	R.Result.Version = version
}
func (R *Redis) Databases() {
	count, err := R.rs.ConfigGet(context.Background(), "databases").Result()
	if err != nil {
		log.Fatalln(err)
	}
	final, err := strconv.Atoi(count["databases"])
	if err != nil {
		log.Fatalln(err)
	}
	R.Result.DataBaseCount = final
	for i := 0; i < final; i++ {
		c := redis.NewClient(
			&redis.Options{Addr: fmt.Sprintf("%s:%s", R.BaseInfo.Host, R.BaseInfo.Port), Password: R.BaseInfo.Password, DB: i},
		)
		res, err := c.Keys(context.Background(), "*").Result()
		if err != nil {
			log.Fatalln(err)
		}
		if len(res) == 0 {
			continue
		}
		values := make(map[string]string, 0)
		for _, v := range res {
			value, _ := c.Get(context.Background(), v).Result()
			values[v] = value
		}
		R.Result.DataBaseInfos = append(R.Result.DataBaseInfos, RedisDatabaseInfo{
			Database:  i,
			KeysCount: len(res),
			Keys:      values,
		})
	}
}
func (R Redis) Reverse() {

}
func (R *Redis) Tables() {

}
