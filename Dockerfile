FROM golang
CP main.go /app/
WORKDIR /app
EXPOSE 80
CMD ["go run", "main.go"]
