FROM golang:latest
WORKDIR /app
COPY . .
RUN go build -o main orders/server/main.go
EXPOSE 80
CMD ["./main"]
