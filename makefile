# Go parameters
VERSION=v1.0.0
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=redis-monitor
BINARY_UNIX=$(BINARY_NAME)_unix

all: build
build:
	$(GOGET) ./...
	$(GOBUILD) -o $(BINARY_NAME) -v
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
	rm -rf redis-monitor-v*
	rm -rf *.sh
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
pack:
	mkdir -p log
	tar zcvf redis-monitor-$(VERSION).tar.gz $(BINARY_NAME) ./conf ./log
dockerpack:
	mkdir -p log
	cp ./docker/* ./
	tar zcvf redis-monitor-$(VERSION).tar.gz $(BINARY_NAME) ./conf *.sh ./log