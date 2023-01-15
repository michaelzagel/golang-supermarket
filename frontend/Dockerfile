FROM golang:latest
WORKDIR /app
COPY . .
RUN go build -o main frontend/server/main.go
EXPOSE 80
CMD ["./main"]
