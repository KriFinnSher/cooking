FROM golang:1.23

WORKDIR /cooking
COPY . .

RUN go mod tidy
RUN go build -o /build ./cmd/server/main.go

EXPOSE 8080
CMD ["/build"]