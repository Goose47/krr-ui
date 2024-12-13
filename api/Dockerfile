FROM golang:alpine AS builder

WORKDIR /build
ADD go.mod .
COPY .. .
RUN go build -o app ./cmd/recs/recs.go

FROM alpine

WORKDIR /app
COPY --from=builder /build/app /app/app
COPY ./config /app/config

CMD ["./app", "--config=./config/local.yml"]