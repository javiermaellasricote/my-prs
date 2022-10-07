FROM golang:alpine AS builder
RUN apk add --no-cache curl git

RUN mkdir /service
ADD . /service
WORKDIR /service
RUN go build -o run .

FROM alpine
RUN mkdir /service
COPY --from=builder /service/run /service/run

ARG RUN_ENV
ENV RUN_ENV ${RUN_ENV}

EXPOSE 8080
CMD ["/service/run"]
