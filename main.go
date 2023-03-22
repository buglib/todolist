package main

import (
	"flag"
	"fmt"
	"todolist/domain"
	"todolist/infra/db"
	"todolist/ui"
)

func main() {
	conf := config()
	fmt.Println(conf)

	err := db.InitMysql(conf.dsn)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&domain.Task{})
	router := ui.InitRouter()
	addr := fmt.Sprintf("%s:%d", conf.host, conf.port)
	router.Run(addr)
}

type Config struct {
	host string
	port int
	dsn  string
}

func config() Config {
	var (
		host      string
		port      int
		mysqlHost string
		mysqlPort int
		user      string
		passwd    string
		db        string
	)

	flag.StringVar(
		&host,
		"host",
		"0.0.0.0",
		"host name or IP address, default is '0.0.0.0'",
	)

	flag.IntVar(
		&port,
		"port",
		8080,
		"port of the application, default is '8080'",
	)

	flag.StringVar(
		&mysqlHost,
		"mysqlHost",
		"0.0.0.0",
		"host name or IP adress of mysql instance",
	)

	flag.IntVar(
		&mysqlPort,
		"mysqlPort",
		3306,
		"port of mysql instance",
	)

	flag.StringVar(
		&user,
		"userName",
		"buglib",
		"user name of mysql",
	)

	flag.StringVar(
		&passwd,
		"passwd",
		"123456",
		"password of mysql user",
	)

	flag.StringVar(
		&db,
		"db",
		"todolist",
		"database name",
	)
	flag.Parse()
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		passwd,
		mysqlHost,
		mysqlPort,
		db,
	)

	config := Config{
		host: host,
		port: port,
		dsn:  dsn,
	}
	return config
}
