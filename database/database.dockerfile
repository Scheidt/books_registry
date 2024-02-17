FROM docker.io/postgres:15.3-alpine

RUN apk update
RUN apk upgrade
RUN apk add bash


EXPOSE 4242