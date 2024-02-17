
# syntax=docker/dockerfile:1

FROM golang:1.21.7-alpine3.19

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY *.go ./
# Build
#RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping .
CMD ["sleep", "infinity"]
RUN go run main.go
EXPOSE 8888

# Run
#CMD ["/docker-gs-ping"]