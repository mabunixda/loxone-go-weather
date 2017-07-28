
CONTAINER=r.nitram.at/loxweather
TODAY=`date +'%Y%m%d'`

container:
	docker build -t ${CONTAINER}:${TODAY} .
	docker tag ${CONTAINER}:${TODAY} ${CONTAINER}:latest
	docker push ${CONTAINER}:latest

golang: goreq
	export GOPATH=${PWD}
	go build -o loxonegoweather

goreq:
	export GOPATH=${PWD}
	go get github.com/Sirupsen/logrus

all: container

