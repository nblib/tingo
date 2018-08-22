package tmsql

import (
	"errors"
	"github.com/nblib/tingo/tdatabase/tdbcomn"
)

// 存放配置的db连接
var dbs map[string]*tdbcomn.DBOperator

/**
通过指定名称获取db操作,名称必须和配置文件中相同,区分大小写
*/
func Get(name string) *tdbcomn.DBOperator {
	if operator, ok := dbs[name]; ok {
		return operator
	}
	panic(errors.New("not found db: " + name))
}
