.PHONY: start build
build:
	go build -o ./urlzip ./cmd/

clean:
	rm ./urlzip

start: build
	./urlzip