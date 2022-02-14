FROM golang:alpine3.15 AS builder

# Install Tools and dependencies
RUN apk add --update --no-cache openssl-dev musl-dev zlib-dev curl tzdata

WORKDIR /build

COPY . .

RUN go build -o out/bin/seating seating.go 


# build final container
FROM alpine

COPY --from=builder /build/out/bin/seating .
#COPY --from=builder /build/.env .

ENTRYPOINT ["/seating"]

CMD ["/seating"]