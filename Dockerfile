FROM golang:1.19.7

ENV APP_PATH=/go/src/github.com/rudderstack_source_app

COPY . $APP_PATH

WORKDIR $APP_PATH

RUN go build -o source-app

EXPOSE 8080

CMD ["./source-app"]

