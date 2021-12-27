FROM golang:1.17.5 AS build

ENV GOPATH /go
WORKDIR /go/src/dew-server

COPY go.mod go.sum /go/src/dew-server/
RUN go mod download

COPY . .
RUN go build -o dew-server

FROM gcr.io/distroless/base-debian11

COPY --from=build /go/src/dew-server/dew-server /bin/dew-server

EXPOSE 8080/tcp
ENTRYPOINT ["dew-server"]
