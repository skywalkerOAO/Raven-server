package dao

import "github.com/skywalkerOAO/gosql"

func init() {
	gosql.DBRegister("redis", "127.0.0.1", "", "", "", 16379, "rds")
	gosql.DBRegister("mssql", "127.0.0.1\\MYNEWSQL", "plan", "sa", "W991224z", 1433, "ms")
}
