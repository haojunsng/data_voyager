FROM --platform=linux/amd64 golang:1.23 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Build the producer
RUN go build -o /weather-producer ./producer

# Build the consumer
RUN go build -o /weather-consumer ./consumer

# Default command to run the binary - producer or consumer
CMD ["go", "version"]
