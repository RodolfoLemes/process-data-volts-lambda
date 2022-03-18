build:
	GOOS=linux go build -o bin/main
	rm -r -f process-data-volts-lambda
	zip process-data-volts-lambda bin/main