# firl.us URL shortener
firl.us provides a small http server in go which provides a way to shorten URLs with your own domain as well as an API and an admin panel to manage shortcuts.

I started this project mainly to learn the Go programming language and ecosystem.
Feel free to open issues and pull requests!

## Installation
Just run src/firl.us/firl.us.go

### API Documentation

### POST /api/shortcut/
200 If inserted
400 If params are wrong
    - path len > 50
    - url is not an url
    - param is missing
409 If already exists
500 If another error occurs

### GET /api/shortcuts/:path
200 If success {"path", "url"}
404 If not exists
500 If another error occurrs

### PUT /api/shortcuts/:path
200 If success
400 If params are wrong
404 If not exists
500 If another error occurrs

### DELETE /api/shortcuts/:path
200 If deleted
404 If not exists
500 If another error occurrs
