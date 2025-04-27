FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o quotes ./main.go

CMD [ "./quotes" ]

FROM builder

