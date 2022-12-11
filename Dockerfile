FROM golang:1.19-alpine

ENV CONNECTION_URI mongodb+srv://test@localhost

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY / ./

RUN go build -o ./mongo-go ./cmd 
EXPOSE 8080

CMD ["./mongo-go"]
