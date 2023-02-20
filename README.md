## Getting started
### Initial setup
Use below Dockerfile (there is no `go.mod` file at this time).

```Dockerfile
FROM golang:1.17.1-alpine

ENV ROOT=/go/src/app
ENV CGO_ENABLED 0
WORKDIR ${ROOT}

RUN apk update && apk add git
EXPOSE 8080
```

Open in container and run the following.

```shell
$ go mod init
$ go install github.com/gin-gonic/gin@v1.7.4
$ go install golang.org/x/tools/gopls@v0.7.2
$ go install github.com/ramya-rao-a/go-outline@latest
$ go install github.com/jstemmer/go-junit-report@v0.9.1
$ go install golang.org/x/tools/cmd/godoc@v0.1.7
$ go install honnef.co/go/tools/cmd/staticcheck@2020.2.1
```

This will create `go.mod` file. 

### Developing
After initial setup you can open this repository in below Dockerfile.

```Dockerfile
FROM golang:1.17.1-alpine

ENV ROOT=/go/src/app
ENV CGO_ENABLED 0
WORKDIR ${ROOT}

RUN apk update && apk add git
COPY go.mod go.sum ./
RUN go mod download
EXPOSE 8080
```

### Local testing
Run the following:

```shell
$ go test ./test -v 2>&1 | go-junit-report > report.xml
```

### Running server

```shell
$ go build -t gin-sample .
$ docker run --name gin-app -it --rm  --publish 0.0.0.0:8080:8080 gin-sample
```

### godoc
Run the following:

```shell
$ go doc <package name>.<function name>
# example
$ go doc controllers.CalcCollatz
package controllers // import "app/controllers"

func CalcCollatz(n int64)
```

### Folder structure
```shell
controllers
    |- collatz.go # Define returning collatz.html
    |- display_collatz.go # Calculate Collatz sequence and display it
    |- healthy.go # Define healthcheck
    |- index.go # Define returning index.html
scripts # Scripts for AWS CodeDeploy hooks
server
    |- view # HTML template files
    |- router.go # Define routing
    |- server.go # Run router
test # Test codes
appspec.yml # AWS CodeDeploy appspec file
buildspec.yml # AWS CodeBuild buildspec file
go.mod # Go modules
main.go
```