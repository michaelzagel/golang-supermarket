FROM golang:latest
WORKDIR /app
COPY . .
RUN go build -o main oldfrontend/server/main.go
EXPOSE 80
CMD ["./main"]
