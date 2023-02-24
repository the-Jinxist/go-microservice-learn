# base go image
FROM golang:1.18-alpine AS builder

RUN mkdir /app

COPY . /app
WORKDIR /app

#This line tells docker that we are not using any C libraries, just the standart libraries
RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api

RUN  chmod +x /app/brokerApp

#build a tiny docker image
FROM alpine:latest

#creates the app directory
RUN mkdir /app

#copies the executable created from the builder step to the /app directory just created
COPY --from=builder /app/brokerApp /app

#provies an entry point for the executable file
CMD [ "/app/brokerApp" ]