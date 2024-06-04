build:
	CGO_ENABLED=0 go build -o ./out/gosite ./src/main.go && podman build . -t gosite
