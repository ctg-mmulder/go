# go

The game Go, in Go.

## To Go

```bash
go run .
```

## Useful commands

```
go mod init <module-name>
go mod edit -replace example.com/greetings=../greetings # Replace a module with a local version
go mod tidy
go test -v # Runs all files ending in _test.go
go build
go list -f '{{.Target}}'
go install
```
