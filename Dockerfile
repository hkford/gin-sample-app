FROM golang:1.20.1-alpine AS builder

ENV ROOT=/app
ENV CGO_ENABLED 0
WORKDIR ${ROOT}

RUN apk update && apk add git
COPY . .
RUN go mod download
RUN go build -o webapp .


FROM gcr.io/distroless/static:nonroot
COPY --from=builder --chown=nonroot:nonroot /app/webapp ./webapp
EXPOSE 8080
ENTRYPOINT [ "./webapp"]