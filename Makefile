main: build

build:
	go build -o main main.go scan.go

run:
	go run main.go
