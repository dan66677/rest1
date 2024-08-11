FROM golang

WORKDIR /app

COPY film/ .

RUN go mod download

RUN go build -o main .

CMD ["./main"]