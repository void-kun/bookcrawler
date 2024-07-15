FROM golang:1.10-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY crawl/ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /bws

EXPOSE 8080

# Execute
CMD ["./bws"]