# vim: set softtabstop=2 shiftwidth=2:
SHELL = bash


CURDIR = ${PWD}

all: install start

install:
	@echo "Build server" 
	$(shell rm ./bin/index)
	$(shell export GOPATH=${CURDIR}; go build -o ./bin/index ./src/index.go)

start:
	@echo "Start server" 
	./bin/index -img ${CURDIR}/html/

