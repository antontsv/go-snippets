Echo server
===========

Example of the basic http server:

Start server:
`go run echo.go 127.0.0.1:4040`

Client requests:

`curl -i -d'{"hello":1}' http://127.0.0.1:4040`

`curl -i http://127.0.0.1:4040`