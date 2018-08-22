package tmsql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nblib/tingo/tdatabase/tdbcomn"
	"time"
)

const (
	SQL_DRIVER  = "mysql"
	URL_PATTERN = "%s:%s@tcp(%s):%d"
)

func init() {
	dbConfs, err := tdbcomn.LoadDBConfig(SQL_DRIVER)
	if err != nil {
		panic(err)
	}
	dbs = make(map[string]*tdbcomn.DBOperator, len(dbConfs))
	for k, conf := range dbConfs {
		db, err := sql.Open(SQL_DRIVER, fmt.Sprintf(URL_PATTERN, conf.User, conf.Password, conf.Host, conf.Port))
		if err != nil {
			panic(err)
		}
		db.SetConnMaxLifetime(time.Duration(conf.MaxLifeTime) * time.Millisecond)
		db.SetMaxIdleConns(conf.MaxIdleConns)
		db.SetMaxOpenConns(conf.MaxOpenConns)

		dbs[k] = tdbcomn.NewOperator(db)
	}
}
