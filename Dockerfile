FROM golang:1.23.3-alpine3.20 as builder

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o hivebox .

FROM alpine:3.20

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/hivebox /usr/local/bin/hivebox

EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/hivebox"]
