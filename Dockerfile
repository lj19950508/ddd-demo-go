FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .
COPY . .
RUN go build -o hello hello.go


FROM alpine

WORKDIR /build
COPY --from=builder /build/hello /build/hello

CMD ["./hello"]

