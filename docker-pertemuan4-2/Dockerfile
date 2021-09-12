FROM golang:alpine
WORKDIR /app
COPY . .
ENV NAMA=FreslyLS
ENV PORT=8000
EXPOSE 8000
RUN go build
CMD ["./docker-p5"]
