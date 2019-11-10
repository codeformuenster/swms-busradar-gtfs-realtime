#FROM golang:1.13-alpine AS build
FROM golang:1-alpine AS build

WORKDIR /gtfs-rt/

RUN apk add git ca-certificates && update-ca-certificates

COPY . .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /main

FROM scratch AS final

COPY --from=build /main /usr/local/bin/main
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT [ "main" ]
