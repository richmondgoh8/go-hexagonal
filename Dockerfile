# Start From Golang Alpine Base Image
FROM golang:1.19.6-alpine as base

# Download Git, Bash
RUN apk update && apk upgrade && apk add --no-cache bash git

# Uncomment Line 8 & 9 if not needed
COPY personal.crt /usr/local/share/ca-certificates/personal.crt
RUN chmod 644 /usr/local/share/ca-certificates/personal.crt && update-ca-certificates

RUN apk add --no-cache git ca-certificates && update-ca-certificates


FROM base as build-env

RUN mkdir /app
WORKDIR /app
# Copy from source from current directory to working directory
COPY . .
#RUN echo $(ls -1 /app)
RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o demo ./cmd/server.go

FROM scratch

COPY --from=build-env /app/demo .
COPY dev.env .

EXPOSE 8080
ENTRYPOINT ["./demo", "dev"]