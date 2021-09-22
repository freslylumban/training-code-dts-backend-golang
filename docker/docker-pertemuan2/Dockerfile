FROM alpine
RUN apk add vim
RUN apk add unzip
WORKDIR /app
RUN echo "test lagi dockerfile" >> automate.txt
COPY . .
CMD ["cat", "/app/automate.txt"]