AUTH_BINARY_NAME=czwrMailingAuth

swag:
	swag init -d ./cmd/auth/ -o ./doc -g main.go --parseDependency

build:
	go build -o bin/${AUTH_BINARY_NAME}.exe ./cmd/auth/main.go

run:
	export SIGNING_KEY=MYSERCRETKEY
	bin/${AUTH_BINARY_NAME}.exe --host 0.0.0.0 --port 5000

build_and_run: swag build run
	echo "build_and_run"

build_docker:
	docker build --tag czwr-mailing-auth .
	docker push chipazawra