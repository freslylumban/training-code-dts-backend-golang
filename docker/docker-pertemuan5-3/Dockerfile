FROM golang:alpine
WORKDIR /app
COPY . .
ENV PORT=8000
EXPOSE 8000
RUN go build -o /app/docker-bisa
CMD ["/app/docker-bisa"]
