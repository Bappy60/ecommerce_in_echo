# Build stage
FROM golang:1.20.3-alpine AS builder
ENV GO111MODULE=on

RUN apk add git
RUN mkdir /app
WORKDIR /app
ADD . .


COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
COPY app.env .
RUN go build -o /app/bin/ecommerce_api

# Final stage
FROM alpine:3.14
COPY --from=builder /app/bin/ecommerce_api /app/bin/ecommerce_api
EXPOSE 9011
CMD [ "/app/bin/ecommerce_api" ]
