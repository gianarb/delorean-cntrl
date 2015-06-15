all: deps

deps:
	go get git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git

run-pub:
	go run ./main.go

build:
	go build
