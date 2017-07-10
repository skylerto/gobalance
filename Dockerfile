FROM golang:1.8.3-alpine

RUN mkdir /src
RUN mkdir /app
COPY . /src

WORKDIR /src
RUN go build -o run;
RUN cp /src/run /app


FROM alpine:latest

WORKDIR /app/
COPY --from=0 /app .

EXPOSE 8080
ENTRYPOINT ["./run"]
