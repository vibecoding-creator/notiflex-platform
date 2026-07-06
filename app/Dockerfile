FROM golang:1.25-alpine AS builder

WORKDIR /build
COPY go.mod .
COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -o api main.go

FROM scratch

COPY --from=builder /build/api /api

EXPOSE 8080

ENTRYPOINT ["/api"]
