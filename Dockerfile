FROM golang:1.22

WORKDIR /go-library

COPY . .

RUN go build -o myapp .

EXPOSE 8080

CMD ["./myapp"]
