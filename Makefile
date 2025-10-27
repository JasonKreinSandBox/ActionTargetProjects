BINARY_NAME=main-linux
build:
	SET GOOS=linux
	SET GOARCH=amd64
	go mod init github.com/JasoKreinSandBox/ActionTargetProjects/main
	go build -o $(BINARY_NAME)
