# ARM § run
FROM ubuntu:latest

# Install go
RUN apt-get update && apt-get install -y golang git

ADD . /app

RUN export CGO_ENABLED=0
RUN export GOOS=linux
RUN export GOARCH=arm64

WORKDIR /app

RUN go build -o ./bin/blockchain-linux-arm64

# Export blockchain.exe to host
VOLUME /app/bin

ENTRYPOINT ["/bin/bash"]