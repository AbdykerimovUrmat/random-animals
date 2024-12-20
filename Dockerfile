# syntax=docker/dockerfile:1

FROM golang:1.22.6

# Set destination for COPY
WORKDIR /app

COPY . .

RUN go mod download

EXPOSE 8080

# Run
CMD ["go", "run",  "./main.go"]