.PHONY: build install

build:
	go build -o totp main.go 

install: build
	mv totp /usr/local/bin