package conf

const DriverName = "mysql"

type DbConf struct {
	Host string
	Port int
	User string
	Pwd string
	DbName string
}

var MasterDbConfig DbConf = DbConf{
	Host:   "localhost",
	Port:   3306,
	User:   "root",
	Pwd:    "123456",
	DbName: "web_forum",
}

var SlaveDbConfig DbConf = DbConf{
	Host:   "localhost",
	Port:   3306,
	User:   "root",
	Pwd:    "123456",
	DbName: "web_forum",
}
