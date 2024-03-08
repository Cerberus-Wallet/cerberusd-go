FROM golang:1.18

RUN mkdir /cerberusd-go
WORKDIR /cerberusd-go
COPY . /cerberusd-go

RUN apt-get update
RUN apt-get install -y redir

RUN go build .

ENTRYPOINT '/cerberusd-go/scripts/run_in_docker.sh'
EXPOSE 11325
