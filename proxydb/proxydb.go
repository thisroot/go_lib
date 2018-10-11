package proxydb

import "database/sql"
import "virt88.aetp.nn/doc2/go_lib/promclient"
import "time"

type ProxyDB struct {
	DB *sql.DB
}

func (self *ProxyDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	startTime := time.Now()
	rows, err := self.DB.Query(query)
	if err != nil {
		promclient.IncError("sql_fetch_rows")
	}
	promclient.EndDbRequestTimer(query, startTime)
	return rows, err
}

func (self *ProxyDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	startTime := time.Now()
	result, err := self.DB.Exec(query)
	promclient.EndDbRequestTimer(query, startTime)
	if err != nil {
		promclient.IncError("sql_fetch_rows")
	}
	return result, err
}
