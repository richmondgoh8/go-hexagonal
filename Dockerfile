# Start From Golang Alpine Base Image
FROM golang:1.19.6-alpine

# Download Git, Bash
RUN apk update && apk upgrade && apk add --no-cache bash git

# Uncomment Line 8 & 9 if not needed
COPY personal.crt /usr/local/share/ca-certificates/personal.crt
RUN chmod 644 /usr/local/share/ca-certificates/personal.crt && update-ca-certificates

RUN apk add --no-cache git ca-certificates && update-ca-certificates
WORKDIR /app

# Copy from source from current directory to working directory
COPY . .
#RUN echo $(ls -1 /app)
RUN go mod tidy

RUN go build -o main ./cmd/server.go

EXPOSE 8080

CMD ["./main"]


