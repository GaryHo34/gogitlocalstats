main: build

build:
	go build -o main main.go scan.go commitInfo.go

run:
	go run main.go scan.go commitInfo.go
