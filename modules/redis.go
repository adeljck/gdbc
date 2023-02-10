package modules

import (
	"errors"
	"fmt"
	"gdbc/common"
	"github.com/gomodule/redigo/redis"
	"log"
)

type Redis struct {
	BaseInfo common.DbInfo
	Result   struct {
		Version       string `db:"version"`
		DataBaseInfos []DataBaseInfo
		DataBaseCount int `db:"count"`
	}
}

var conn redis.Conn

func (R *Redis) init() error {
	R.BaseInfo = common.BaseInfo
	dsn := fmt.Sprintf("%s:%s", R.BaseInfo.Host, R.BaseInfo.Port)
	conn, err := redis.Dial("tcp", dsn, redis.DialPassword(R.BaseInfo.Password))
	//defer client.Close()
	if err != nil {
		return err
	}
	reply, err := conn.Do("ping")
	if err != nil {
		return err
	}
	if reply == "PONG" {
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
	result, err := conn.Do("info")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(result)
}
func (R *Redis) Databases() {

}
func (R Redis) Reverse() {

}
func (R *Redis) Tables() {

}
