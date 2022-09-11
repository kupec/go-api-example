FROM golang:1.19-alpine as builder
RUN apk update && apk add build-base

WORKDIR /build
COPY go.mod go.sum .
RUN go mod download

COPY . .
RUN go build -o app

FROM scratch
COPY --from=builder /build/app .
ENTRYPOINT ["./app"]

