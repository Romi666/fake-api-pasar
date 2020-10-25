FROM golang:alpine

ADD . /app

WORKDIR /app

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

RUN chmod +x ./fake-web-pasar-api

EXPOSE 80

CMD /app/fake-web-pasar-api