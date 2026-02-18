FROM golang:1.26

WORKDIR /app

COPY math.go .
COPY math_test.go .

RUN go mod init mymath

RUN go build -o math

CMD ["./math"]
