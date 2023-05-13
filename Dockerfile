FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . ./
RUN go mod tidy

RUN go build -o main main.go

FROM scratch

COPY --from=builder /build/main .

EXPOSE 5000

ENTRYPOINT ["./main"]