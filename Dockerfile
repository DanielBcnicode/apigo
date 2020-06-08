FROM golang:alpine as builder

LABEL maintainer="Daniel Martinez <daniel@bcnicode.com>"

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download 

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .
#COPY --from=builder /app/.env .       

EXPOSE 8000

CMD ["./main"]