AUTH_BINARY_NAME=czwr-mailing-auth

swag:
	swag init -d ./cmd/auth/ -o ./doc -g main.go --parseDependency

build:
	go build -o bin/${AUTH_BINARY_NAME}.exe ./cmd/auth/main.go

run:
	export SIGNING_KEY=MYSECRETKEY
	SIGNING_KEY=MYSECRETKEY
	bin/${AUTH_BINARY_NAME}.exe --host 0.0.0.0 --port 8885

build_and_run: swag build run
	echo "build_and_run"

build_docker:
	docker build --tag czwr-mailing-auth .