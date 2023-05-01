visualgit: build

build:
	go build -o visualgit main.go scan.go commitInfo.go

run:
	go run main.go scan.go commitInfo.go
