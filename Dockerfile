FROM golang:1.21.4-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=O GOOS=linux go build -a -installsufix nocgo -o myapp cmd/api/main.go .

FROM scratch

COPY --from=builder /app/myapp /myapp

EXPOSE $PORT

CMD ["/myapp"]
