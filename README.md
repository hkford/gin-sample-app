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
go mod init
go install github.com/gin-gonic/gin@v1.7.4
go install golang.org/x/tools/gopls@v0.7.2
go install github.com/ramya-rao-a/go-outline@latest
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