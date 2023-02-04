package run

import (
	"dbconnector/modules"
)

func Run() {
	m := modules.Mysql{}
	modules.Checker(m)
	modules.GetVersion(m)
}
