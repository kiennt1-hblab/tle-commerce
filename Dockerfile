FROM golang:1.21-alpine AS builder

RUN apk add --no-cache gcc musl-dev sqlite-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o bookstore-server cmd/server/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates sqlite-libs

WORKDIR /root/

COPY --from=builder /app/bookstore-server .
COPY --from=builder /app/migrations ./migrations

RUN mkdir -p data

EXPOSE 3000

ENV PORT=3000
ENV DB_PATH=./data/bookstore.db
ENV ALLOWED_ORIGINS=*

CMD ["./bookstore-server"]
