ARG GO_VERSION=1.21.0
FROM golang:${GO_VERSION}-bookworm as builder

RUN mkdir -p /Users/aa/bucket
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
ENV DATABASE_URL=/Users/aa/private/script
RUN go build -v -o /usr/src/app/run-app .

FROM debian:bookworm
RUN apt-get update
RUN apt-get install -y vim jq wget curl openjdk-17-jre-headless htop procps rsyslog bash-completion python3 python3-pip python3-selenium python3-bs4 python3-parsel unzip gzip

RUN rm -rf /var/lib/apt/lists/*
EXPOSE 8080

COPY --from=builder /usr/src/app/run-app /usr/local/bin/

CMD ["sh", "-c", "mkdir -p /Users/aa/private & /usr/local/bin/run-app run 8080"]
