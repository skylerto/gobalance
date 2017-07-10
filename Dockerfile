FROM golang:1.8.3-alpine

RUN mkdir /src
RUN mkdir /app
COPY . /src

WORKDIR /src
# RUN apk upgrade && apk add git
RUN apk add --update git
RUN go get github.com/get-ion/ion
RUN go build -o run;
RUN cp /src/run /app


FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /app/
COPY --from=0 /app .

EXPOSE 8080
ENTRYPOINT ["./run"]
