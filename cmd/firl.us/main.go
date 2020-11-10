package main

import (
	"os"
	"strconv"

	"firlus.dev/firl.us/internal/server"

	"firlus.dev/firl.us/internal/store"
)

func main() {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	if port == 0 {
		port = 42001
	}
	//writeAccessPassword := os.Getenv("WRITE_ACCESS_PASSWORD")
	mysqlUser := os.Getenv("MYSQL_USER")
	if mysqlUser == "" {
		mysqlUser = "root"
	}
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	if mysqlPassword == "" {
		mysqlPassword = "bidendid911"
	}
	mysqlServer := os.Getenv("MYSQL_SERVER")
	if mysqlServer == "" {
		mysqlServer = "localhost:3306"
	}
	mysqlDbName := os.Getenv("MYSQL_DBNAME")
	if mysqlDbName == "" {
		mysqlDbName = "firlus_url"
	}
	store := store.Setup(mysqlUser, mysqlPassword, mysqlServer, mysqlDbName)
	server.Setup(port, store)
}
