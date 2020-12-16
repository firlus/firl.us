# 🌐 firl.us URL shortener

firl.us provides a small http server in go which provides a way to shorten URLs with your own domain as well as an API and an admin panel to manage shortcuts.

I started this project mainly to learn the Go programming language and ecosystem.
Feel free to open issues and pull requests!

## Installation

### 🏗️ Build it yourself

Make sure Go, Node.js and NPM are installed to build the application. To build the server run `go build cmd/firl.us/firl.us.go`. To build the admin panel, switch to the `web` folder and run `npm i` followed by `npm run build`.

Now copy the `web/public` folder and the server binary into the same folder, to make sure the binary finds the website to serve.
(I will try to automate this process with a Makefile in the future...)

You also need a MySQL-Database to store the shortcuts.

#### ⌨️ Command line arguments

- `-port`: Port on which you can access admin panel and API (defaults to `80`)
- `-mysql-server`: Address of database server (defaults to `localhost:3306`)
- `-mysql-user`: MySQL username (defaults to `root`)
- `-mysql-password`: MySQL password (defaults to `mysqlpassword`)
- `-mysql-db`: MySQL database name (defaults to `firlus_url`; must already exist on MySQL server)
- `-password`: Admin panel password (defaults to `admin`)

### 🐋 Docker

I provide Docker images for this application: [firlus/firl.us](https://hub.docker.com/firlus/firl.us).

The provided containers expose port `80`.

To configure database and admin password, pass the following environment variables to the container:

- `MYSQL_SERVER`
- `MYSQL_USER`
- `MYSQL_PASSWORD`
- `MYSQL_DB`
- `ADMIN_PASSWORD`

> Please refer to the command line arguments section for the usage and defaults of the variables.

#### 🐙 Docker Compose

If you want a ready to use instance, use the provided `docker-compose.yml`. Do not forget to change the admin password inside of this file, otherwise it will be `admin`, which is not very secure.

Please ignore all the traefik-related labels (comment them out if you use traefik yourself), they are for my own configuration and won't work on your system, since the domain `firl.us` belongs to me.

## API Documentation

There is more documentation to do...

### POST /api/shortcut/

- 200 If inserted
- 400 If params are wrong
- 409 If already exists
- 500 If another error occurs

### GET /api/shortcuts/:path

- 200 If success {"path", "url"}
- 404 If not exists
- 500 If another error occurrs

### PUT /api/shortcuts/:path

- 200 If success
- 400 If params are wrong
- 404 If not exists
- 500 If another error occurrs

### DELETE /api/shortcuts/:path

- 200 If deleted
- 404 If not exists
- 500 If another error occurrs
