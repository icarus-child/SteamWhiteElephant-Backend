FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o server

FROM alpine
WORKDIR /app
ENV GIN_MODE=release
COPY --from=builder /app/server .
EXPOSE 3333
CMD ["./server"]
