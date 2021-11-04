FROM golang:1.15.4

ENV PORT 8080
ENV APP_NAME myproject

COPY . /go/src/${APP_NAME}
COPY ./connection /go/src/${APP_NAME}/connection
COPY ./customhttp /go/src/${APP_NAME}/customhttp
COPY ./dataacces /go/src/${APP_NAME}/database

WORKDIR /go/src/${APP_NAME}

RUN go build

CMD ./${APP_NAME}

EXPOSE ${PORT}