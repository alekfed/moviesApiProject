# Start by building the application.
FROM golang:1.17.7-alpine3.15 as build

WORKDIR /go/src/app
ADD . /go/src/app

RUN go get -d -v ./...

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -extldflags '-static'" -o /go/bin/app

# Now copy it into our base image.
FROM gcr.io/distroless/base-debian11

COPY --from=build /go/bin/app /

CMD ["/app"]