package db

import (
	"database/sql"
	"io/ioutil"
	"os"
	"path/filepath"

	// mysql is used in database connection, it is slient include
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
)

// Database has database connection
type Database struct {
	Connection *sql.DB
}

var sharedInstance = New()

// New is create Database object with connection pool
func New() *Database {
	env := os.Getenv("APPENV")
	root := os.Getenv("APPROOT")
	path := filepath.Join(root, "backend/config/db.yml")
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(buf, &m)
	if err != nil {
		panic(err)
	}
	username := m[env].(map[interface{}]interface{})["user"].(string)
	password := m[env].(map[interface{}]interface{})["password"].(string)
	database := m[env].(map[interface{}]interface{})["name"].(string)
	host := m[env].(map[interface{}]interface{})["host"].(string)
	pool := m[env].(map[interface{}]interface{})["pool"].(int)
	username = os.ExpandEnv(username)
	password = os.ExpandEnv(password)
	database = os.ExpandEnv(database)
	host = os.ExpandEnv(host)
	db, err := sql.Open("mysql", username+":"+password+"@tcp("+host+":3306)/"+database+"?parseTime=true&charset=utf8")
	if err != nil {
		panic(err)
	}

	// MaxIdle: mysqlへのアクセスがないときにも保持しておくconnection poolの上限
	// MaxOpen: idle + activeなconnection poolの上限数
	db.SetMaxIdleConns(pool)
	db.SetMaxOpenConns(pool)

	return &Database{
		Connection: db,
	}
}

// SharedInstance return database connection object
func SharedInstance() *Database {
	return sharedInstance
}

// Close database connection
func (d *Database) Close() error {
	return d.Connection.Close()
}
