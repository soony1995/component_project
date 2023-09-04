PACKAGENAME := login-pro.exe

build:
	go build -o ${PACKAGENAME}

run:
	./${PACKAGENAME}

restart: build run
