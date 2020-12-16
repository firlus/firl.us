package main

import (
	"flag"

	"firlus.dev/firl.us/internal/server"

	"firlus.dev/firl.us/internal/common"
	"firlus.dev/firl.us/internal/store"
)

// AdminPassword contains password for admin panel
var AdminPassword string

func main() {

	port := flag.String("port", "42001", "Port")
	mysqlUser := flag.String("mysql-user", "root", "User for MySQL-Database")
	mysqlPassword := flag.String("mysql-password", "mysqlpassword", "Password for MySQL-Database")
	mysqlServer := flag.String("mysql-server", "localhost:3306", "Address of MySQL-Server")
	mysqlDbName := flag.String("mysql-db", "firlus_url", "Name of MySQL database")
	adminPassword := flag.String("password", "admin", "Password to access API and admin panel")
	flag.Parse()

	common.SetPassword(*adminPassword)
	store := store.Setup(*mysqlUser, *mysqlPassword, *mysqlServer, *mysqlDbName)
	server.Setup(*port, store)
}
