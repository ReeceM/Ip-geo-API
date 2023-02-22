FROM golang:1.20.1-buster as builder

ENV APP_HOME /app
ENV PORT 8080

WORKDIR $APP_HOME
COPY . $APP_HOME

RUN go mod download
RUN go build -v -o server
RUN CGO_ENABLED=0 go build -o server main.go

FROM alpine:latest
COPY --from=builder $APP_HOME $APP_HOME
EXPOSE $PORT
CMD ["/app/server"]
