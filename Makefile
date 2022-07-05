.PHONY: build
build:
	go build -v ./cmd/apiserver
	go get -u github.com/gorilla/mux
	go get github.com/BurntSushi/toml
	go get github.com/sirupsen/logrus
	go get github.com/stretchr/testify


.DEFAULT_GOAL := build