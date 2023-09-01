PATH := $(PWD)

build:
	go build -o ${PATH}/start

run:
	${PATH}/start

restart:
	go build -o ${PATH}/start
	${PATH}/start

	