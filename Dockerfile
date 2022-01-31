FROM golang:1.15.7-buster

COPY mysql_docker_script.sh /go/src/github.com/pauliusluksys/golang-bankingSystem/
COPY go.mod go.sum /go/src/github.com/pauliusluksys/golang-bankingSystem/
WORKDIR /go/src/gitlab.com/idoko/letterpress
RUN mysql_docker_script.sh
RUN go mod download
COPY . /go/src/github.com/pauliusluksys/golang-bankingSystem
RUN go build -o /usr/bin/golang-bankingSystem ./

EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/golang-bankingSystem"]