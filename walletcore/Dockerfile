FROM golang:1.20

WORKDIR /app/

RUN apt-get update && apt-get install -y librdkafka-dev

CMD ["go", "run", "/app/cmd/walletcore/main.go"]