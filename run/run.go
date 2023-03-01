package run

import (
	"gdbc/modules"
)

func Run() {
	m := modules.CreateObject()
	modules.Checker(m)
	modules.GetVersion(m)
	modules.GetDatabases(m)
	modules.GetTables(m)
	modules.PrintInfo(m)
}
