FROM golang:1.19
LABEL authors="bogumila_walendziak"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o EcoRide .

EXPOSE 8080

CMD ["./EcoRide"]