FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN go get ./web/; CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./web/

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/server .
COPY --from=builder /app/web/assets ./assets
EXPOSE 8080
CMD ["./server"] 
