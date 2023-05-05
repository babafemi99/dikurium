FROM golang:1.19.1-alpine3.16 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o main-app server.go

FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main-app /app
COPY --from=builder /app/app.env /app
EXPOSE 8080
CMD ["/app/main-app"]