package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
)

var MYSQL_HOST = "MYSQL_HOST"
var MYSQL_PORT = "MYSQL_PORT"
var MYSQL_USERNAME = "MYSQL_USERNAME"
var MYSQL_PASSWORD = "MYSQL_PASSWORD"
var MYSQL_DATABASE = "MYSQL_DATABASE"
var MYSQL_TABLE = "MYSQL_TABLE"

func pollMySQL() (bool, error) {
	host := getEnvVar(MYSQL_HOST, "")
	if host == "" {
		return false, fmt.Errorf("MYSQL_HOST is not set")
	}

	port := getEnvVar(MYSQL_PORT, "3306")
	portInt, err := strconv.Atoi(port)
	if err != nil {
		return false, fmt.Errorf("MYSQL_PORT is not a valid integer: %w", err)
	}

	username := getEnvVar(MYSQL_USERNAME, "root")
	if username == "" {
		return false, fmt.Errorf("MYSQL_USERNAME is not set")
	}

	database := getEnvVar(MYSQL_DATABASE, "")
	if database == "" {
		return false, fmt.Errorf("MYSQL_DATABASE is not set")
	}

	password := getEnvVar(MYSQL_PASSWORD, "")
	desiredTable := getEnvVar(MYSQL_TABLE, "")

	config := mysql.NewConfig()
	config.User = username
	config.Passwd = password
	config.Net = "tcp"
	config.Addr = fmt.Sprintf("%s:%d", host, portInt)
	config.DBName = database
	config.Timeout = 2 * time.Second
	config.ReadTimeout = 2 * time.Second

	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return true, err
	}
	defer db.Close()

	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		return true, err
	}

	containsTable := false
	for rows.Next() {
		var tableName string
		err := rows.Scan(&tableName)
		if err != nil {
			return true, err
		}

		if tableName == desiredTable {
			containsTable = true
		}
	}

	if desiredTable != "" && !containsTable {
		return true, fmt.Errorf("table %s does not exist in %s", desiredTable, database)
	}

	fmt.Printf("MySQL is available at %s:%d (database: %s, table: %s)\n", host, portInt, database, desiredTable)
	return false, nil
}
