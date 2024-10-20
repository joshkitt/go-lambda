```
go mod init go-lambda
```
```
GOOS=linux GOARCH=amd64 go build -o handler main.go
```
```
zip handler.zip handler
```


