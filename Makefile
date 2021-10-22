AUTH_BINARY_NAME=czwrMailingAuth

swag:
	swag init -d ./cmd/auth/ -o ./doc -g main.go --parseDependency

build:
	go build -o bin/${AUTH_BINARY_NAME}.exe ./cmd/auth/main.go

run:
	bin/${AUTH_BINARY_NAME}.exe -host 127.0.0.1 -port 5000

build_and_run: swag build run
	echo "build_and_run"
	SIGNING_KEY=MYSERCRETKEY

build_docker:
	docker build --tag czwr-mailing-auth .