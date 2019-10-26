FROM golang:latest

WORKDIR /app

COPY . .

RUN ls

RUN go build seating.go

CMD ["./seating"]