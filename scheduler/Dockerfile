#build stage
FROM golang:1.21.3-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

#Run stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY ./doc/fake_source/ ./doc/fake_source/
COPY repository/migrations ./repository/migrations

EXPOSE 8080
CMD ["/app/main"]