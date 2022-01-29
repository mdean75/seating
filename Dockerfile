FROM golang:alpine3.15 AS builder

# Install Tools and dependencies
RUN apk add --update --no-cache openssl-dev musl-dev zlib-dev curl tzdata

WORKDIR /build

COPY . .

RUN go build -o seating seating.go 



FROM alpine

COPY --from=builder /build/seating .

ENTRYPOINT ["/seating"]

CMD ["/seating"]