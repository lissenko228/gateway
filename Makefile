dev: 
	nodemon --exec go run main.go

start:
	go build main.go

install:
	go get .