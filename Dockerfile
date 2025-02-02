FROM golang:alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o go-app ./cmd/main.go || exit 1

FROM alpine:latest

COPY --from=builder /app/go-app /go-app

EXPOSE 8080

CMD [ "/go-app" ]

