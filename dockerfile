#Builder stage 
FROM golang:1.18.4-alpine AS builder
WORKDIR /app
COPY . .
RUN apk add build-base
RUN go build -tags musl -o main main.go 


#RUN stage
FROM alpine 
WORKDIR /app
COPY --from=builder /app/main .
CMD ["/app/main"]