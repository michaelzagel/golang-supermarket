FROM golang:latest
WORKDIR /app
COPY . .
RUN go build -o main auth/server/main.go
EXPOSE 80
CMD ["./main"]
