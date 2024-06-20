test:
	go build -o test ./*.go
	CONFIG_FILE=./config.yaml ./test
	rm test
testjson:
	go build -o test ./*.go
	CONFIG_FILE=./config.json ./test
	rm test
testFullPath:
	go build -o test ./*.go
	CONFIG_FILE=$$(pwd)/config.yaml ./test
	rm test
testWithFlag:
	go build -o test ./*.go
	CONFIG_FILE=./config.yaml ./test --testflag="test"
	rm test