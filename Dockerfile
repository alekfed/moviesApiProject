# Start by building the application.
FROM golang:1.17.7-alpine3.15 as build

WORKDIR /go/src/app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . /go/src/app

RUN CGO_ENABLED=0 go build -ldflags="-s -w -extldflags '-static'" -o /go/bin/app

# Now copy it into our base image.
FROM gcr.io/distroless/static-debian11

COPY --from=build /go/bin/app /

CMD ["/app"]