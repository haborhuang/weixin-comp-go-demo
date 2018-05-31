PKG="github.com/haborhuang/weixin-comp-go-demo"
IMAGE?=wx-demo
BINARY=service

all:
	go install

clean:
	@rm -fr $(GOPATH)/bin/$(BINARY) bin $(BINARY)

docker-build:
	docker run --rm  -v `pwd`:/go/src/$(PKG) -w /go/src/$(PKG) golang:1.8.3-alpine3.6 go build -o $(BINARY)


build-image: docker-build
	docker build -t $(IMAGE) .

.PHONY: docker-build build-image
