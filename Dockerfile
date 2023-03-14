FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git build-base
RUN apk add -U --no-cache ca-certificates && update-ca-certificates
RUN mkdir /server
WORKDIR /server
COPY ./ ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o build/proton
FROM scratch
LABEL MAINTAINER="Saksham"
LABEL VERSION="0.0.1"
COPY --from=builder /server/build/proton /go/bin/proton
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/go/bin/proton"]
