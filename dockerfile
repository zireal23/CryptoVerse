#Builder stage 
FROM golang:1.18.4-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go


#RUN stage
FROM alpine
WORKDIR /app
COPY --from=builder /app/main .
CMD ["/app/main"]
