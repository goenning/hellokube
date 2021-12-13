FROM golang:1.17-alpine

WORKDIR /app

COPY *.mod ./
COPY *.html ./
COPY *.go ./

RUN GOOS=linux GOARCH=amd64 go build -o ./bin

EXPOSE 8080

CMD [ "/app/bin" ]