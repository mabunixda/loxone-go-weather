
CONTAINER=loxweather
TODAY=`date +'%Y%m%d'`

container: golang
	docker build -t ${CONTAINER}:${TODAY} . 
	docker tag ${CONTAINER}:${TODAY} ${CONTAINER}:latest

golang: goreq
	export GOPATH=${PWD}
	go build -v server.go

goreq:
	export GOPATH=${PWD}
	go get github.com/Sirupsen/logrus

all: container
	
