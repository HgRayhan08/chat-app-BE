#build stage
FROM golang:1.25-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o app .

CMD [ "go", "run", "main.go" ]

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder  /app/app .
# ENTRYPOINT ["/app"]
# LABEL Name=chatappbe Version=0.0.1
EXPOSE 3000
CMD [ "./app" ]