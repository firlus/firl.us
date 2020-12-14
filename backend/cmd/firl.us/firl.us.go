package main

import (
	"flag"
	"fmt"
	"os"

	"firlus.dev/firl.us/internal/server"

	"firlus.dev/firl.us/internal/common"
	"firlus.dev/firl.us/internal/store"
)

// AdminPassword contains password for admin panel
var AdminPassword string

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":80"
	} else {
		port = fmt.Sprintf(":%v", port)
	}
	mysqlUser := os.Getenv("MYSQL_USER")
	if mysqlUser == "" {
		mysqlUser = "root"
	}
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	if mysqlPassword == "" {
		mysqlPassword = "mysqlpassword"
	}
	mysqlServer := os.Getenv("MYSQL_SERVER")
	if mysqlServer == "" {
		mysqlServer = "db:3306"
	}
	mysqlDbName := os.Getenv("MYSQL_DBNAME")
	if mysqlDbName == "" {
		mysqlDbName = "firlus_url"
	}
	adminPassword := flag.String("password", "admin", "Password to access API and admin panel")
	common.SetPassword(*adminPassword)
	store := store.Setup(mysqlUser, mysqlPassword, mysqlServer, mysqlDbName)
	server.Setup(port, store)
}
