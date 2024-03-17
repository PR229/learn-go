FROM golang:1.19


WORKDIR /greetings
COPY ./greetings/go.mod ./greetings/greetings.go ./
RUN go mod download

WORKDIR /

COPY ./sayHello/go.mod ./sayHello/go.sum ./
RUN go mod download
COPY ./sayHello/*.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

EXPOSE 8080

CMD ["/docker-gs-ping"]
