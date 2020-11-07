# buildtime
FROM golang:1.15-buster as build
WORKDIR /go/src/snippetbox
ADD . /go/src/snippetbox
RUN go get -d -v ./...
RUN go build -o /go/bin/snippetbox -v ./cmd/web 

# runtime
FROM gcr.io/distroless/base-debian10
COPY --from=build /go/bin/snippetbox /
CMD ["/snippetbox"]