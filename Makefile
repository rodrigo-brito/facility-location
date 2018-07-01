build:
	go build -i

test: build
	./run.sh
