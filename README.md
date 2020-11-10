# POST /api/shortcuts/
200 If inserted
400 If params are wrong
    - path len > 50
    - url is not an url
    - param is missing
409 If already exists
500 If another error occurs

# GET /api/shortcuts/:path
200 If success {"path","url"}
404 If not exists
500 If another error occurrs

# PUT /api/shortcuts/:path
200 If success
400 If params are wrong
404 If not exists
500 If another error occurrs

# DELETE /api/shortcuts/:path
200 If success
400 If params are wrong
404 If not exists
500 If another error occurrs