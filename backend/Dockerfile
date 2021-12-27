FROM golang:1.17.5 AS build

ENV GOPATH /go
WORKDIR /go/src/dew-backend

COPY go.mod go.sum /go/src/dew-backend/
RUN go mod download

COPY . .
RUN go build -o dew-backend

FROM gcr.io/distroless/base-debian11

COPY --from=build /go/src/dew-backend/dew-backend /bin/dew-backend

EXPOSE 8080/tcp
ENTRYPOINT ["dew-backend"]
