# Go notes

## Useful commands

```
go mod init <module-name>
go mod edit -replace example.com/greetings=../greetings
go mod tidy
go test -v
go build
go list -f '{{.Target}}'
```

## Concepts

-   [Slices](https://go.dev/blog/slices-intro)
-   [Maps](https://go.dev/blog/maps)
-   [Blank identifier](https://go.dev/doc/effective_go#blank)
-   [Range](https://go.dev/wiki/Range)
-   [Tests](https://pkg.go.dev/testing)
    Ending a file's name with \_test.go tells the go test command that this file contains test functions.
-   [Build command](https://pkg.go.dev/cmd/go#hdr-Compile_packages_and_dependencies) compiles the packages, along with their dependencies, but it doesn't install the results.
-   [Install command](https://go.dev/ref/mod#go-install) compiles and installs the packages.
-   [Managing dependencies](https://go.dev/doc/modules/managing-dependencies)
-   [Developing and publishing modules](https://go.dev/doc/modules/developing)

## Tutorial

https://go.dev/doc/tutorial/module-conclusion
