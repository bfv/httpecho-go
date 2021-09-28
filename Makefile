build:
	GOOS=linux GOARCH=amd64 go build -o build/httpecho-linux .
	GOOS=windows GOARCH=amd64 go build -o build/httpecho-win64 .
