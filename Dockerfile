FROM golang:1.14.6-alpine3.12 as builder

COPY go.mod go.sum /go/src/github.com/kyler-swanson/truncr/

WORKDIR /go/src/github.com/kyler-swanson/truncr

RUN go mod download

COPY . /go/src/github.com/kyler-swanson/truncr

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/truncr github.com/kyler-swanson/truncr

FROM alpine

RUN apk add --no-cache ca-certificates && update-ca-certificates

COPY --from=builder /go/src/github.com/kyler-swanson/truncr/build/truncr /usr/bin/truncr

EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/truncr"]