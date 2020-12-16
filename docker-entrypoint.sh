#!/bin/sh

# Make sure all env variables are set

if [ -z "$PORT" ]
then
      PORT="80"
fi

if [ -z "$MYSQL_SERVER" ]
then
      MYSQL_SERVER="localhost:3306"
fi

if [ -z "$MYSQL_USER" ]
then
      MYSQL_USER="localhost:3306"
fi

if [ -z "$MYSQL_PASSWORD" ]
then
      MYSQL_USER="mysqlpassword"
fi

if [ -z "$MYSQL_DB" ]
then
      MYSQL_DB="firlus_url"
fi

if [ -z "$MYSQL_PASSWORD" ]
then
      MYSQL_DB="admin"
fi

./firlus-url -port $PORT -mysql-server $MYSQL_SERVER -mysql-user $MYSQL_USER -mysql-password $MYSQL_PASSWORD -mysql-db $MYSQL_DB -password $ADMIN_PASSWORD