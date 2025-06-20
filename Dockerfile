FROM golang:1.24-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.* ./

RUN go mod download
RUN go mod tidy

COPY . .

# RUN go build -o main ./main.go

EXPOSE 5000

# CMD [ "./main" ]
CMD ["air", "-c", ".air.toml"]