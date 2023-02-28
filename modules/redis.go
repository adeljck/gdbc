package modules

import (
	"context"
	"errors"
	"fmt"
	"gdbc/common"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
	"strings"
)

type Redis struct {
	BaseInfo common.DbInfo
	Result   struct {
		Version       string `db:"version"`
		DataBaseInfos []DataBaseInfo
		DataBaseCount int `db:"count"`
	}
}

var client *redis.Client

func (R *Redis) init() error {
	R.BaseInfo = common.BaseInfo
	dsn := fmt.Sprintf("%s:%s", R.BaseInfo.Host, R.BaseInfo.Port)
	client = redis.NewClient(
		&redis.Options{Addr: dsn, Password: R.BaseInfo.Password, DB: 0},
	)
	pong, err := client.Ping(context.Background()).Result()
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
func (R *Redis) Info() {
	fmt.Println("********************************************************************************************")
	results := "Database Type: %s\nDatabase Version: %s\nHost: %s\nPort: %s\nUser: %s\nPassword: %s\n"
	fmt.Printf(results, R.BaseInfo.DbType, R.Result.Version, R.BaseInfo.Host, R.BaseInfo.Port, R.BaseInfo.UserName, R.BaseInfo.Password)
}
func (R *Redis) Version() {
	info, err := client.Info(context.Background(), "server").Result()
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
	count, err := client.ConfigGet(context.Background(), "databases").Result()
	if err != nil {
		log.Fatalln(err)
	}
	final, err := strconv.Atoi(count["databases"])
	if err != nil {
		log.Fatalln(err)
	}
	R.Result.DataBaseCount = final
	for i := 0; i < final; i++ {
		tmp := DataBaseInfo{
			Database:   string(i),
			TableCount: 0,
			Tables:     nil,
		}
		R.Result.DataBaseInfos = append(R.Result.DataBaseInfos, tmp)
	}
}
func (R Redis) Reverse() {

}
func (R *Redis) Tables() {
	for _, v := range R.Result.DataBaseInfos {
		client.
	}
}
