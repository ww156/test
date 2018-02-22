FROM golang
CP main.go /app/
WORKDIR /app
EXPOSE 8080
CMD ["go run", "main.go"]
