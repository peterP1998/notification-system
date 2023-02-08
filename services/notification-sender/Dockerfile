# Stage 1: compile the program
FROM golang:1.20 as build-stage
WORKDIR /app
COPY . .
RUN go mod download
COPY . .
RUN go build -o notification-sender main.go

# Stage 2: build the image
FROM alpine:latest  
RUN apk --no-cache add ca-certificates libc6-compat
WORKDIR /app/
COPY --from=build-stage /app .
EXPOSE 8021
CMD ["./notification-sender"]  