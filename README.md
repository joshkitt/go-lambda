```
go mod init go-lambda
```

```
GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o bootstrap main.go
```

```
zip handler.zip handler
```
