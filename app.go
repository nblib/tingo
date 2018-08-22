package main

import (
	"github.com/nblib/tingo/tdatabase/tmsql"
)

func main() {
	operator := tmsql.Get("userWriteDB")
	operator.QueryRow("select * from test")
}
