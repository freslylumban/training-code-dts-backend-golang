FROM golang:alpine
RUN apk add vim
WORKDIR /app
COPY . .
CMD ["go", "run", "/app/main.go"]
