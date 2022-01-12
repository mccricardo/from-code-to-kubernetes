FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /server

EXPOSE 8090

CMD [ "/server" ]
