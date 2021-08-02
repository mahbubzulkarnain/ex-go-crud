package config

import (
	"fmt"
	"time"

	"github.com/gomodul/envy"
)

type databaseItem struct {
	DriverName      string
	DataSourceName  string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

type database struct {
	MySQL databaseItem
}

var DB = database{
	MySQL: databaseItem{
		DriverName: "mysql",
		DataSourceName: fmt.Sprintf(
			"%v:%v@tcp(%v:%v)/%v?parseTime=true&charset=utf8",
			envy.Get("MYSQL_USER"),
			envy.Get("MYSQL_PASS"),
			envy.Get("MYSQL_HOST", "127.0.0.1"),
			envy.Get("MYSQL_PORT", "3306"),
			envy.Get("MYSQL_NAME", "todo"),
		),
		MaxIdleConns:    envy.GetInt("MYSQL_MAX_IDLE_CONNS", "10"),
		MaxOpenConns:    envy.GetInt("MYSQL_MAX_OPEN_CONNS", "100"),
		ConnMaxLifetime: envy.GetDuration("MYSQL_MAX_LIVE_TIME_IN_SECOND", "3600") * time.Second,
	},
}
