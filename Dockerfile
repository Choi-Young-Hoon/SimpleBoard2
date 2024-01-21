FROM ubuntu:latest

RUN apt-get update && \
    apt-get install -y wget && \
    apt-get install -y build-essential

RUN wget https://go.dev/dl/go1.21.6.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.21.6.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin

WORKDIR /app

COPY . .

RUN go build -o myapp

CMD ["./myapp"]
