# vim: set softtabstop=2 shiftwidth=2:
SHELL = bash


CURDIR = ${PWD}
export GOPATH=${CURDIR}

all: clean install start

clean:
	@echo "Clean server" 
	$(shell rm ./bin/index)


install:
	@echo "Build server" 
	go build -o ./bin/index ./src/index.go


start:
	@echo "Start server" 
	./bin/index -dir ${CURDIR} -port 19720

test:
	@echo "test server"
	go get gopkg.in/check.v1 
	go test ./src/game/
