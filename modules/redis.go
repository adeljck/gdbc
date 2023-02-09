package modules

import "gdbc/common"

type Redis struct {
	BaseInfo common.DbInfo
	Result   struct {
		Version       string `db:"version"`
		DataBaseInfos []DataBaseInfo
		DataBaseCount int `db:"count"`
	}
}
