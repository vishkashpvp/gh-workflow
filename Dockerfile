FROM golang:1.20

WORKDIR /app

COPY server/go.mod server/go.sum ./

RUN go mod download && go mod verify

COPY server/ .

RUN go build -o main .

EXPOSE 8080

CMD [ "./main" ]
