
build:
	go build .

image:

dep:
	go mod tidy
	go mod vendor
