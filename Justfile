compile:
    @printf "\033[0;33mCompiling proto files for server[Language: Golang] \033[0m \n"
    @protoc --go_out=./server --go-grpc_out=./server proto/*.proto
    @printf "\033[0;33mCompiling proto files for cli[Language: Golang] \033[0m \n"
    @protoc --go_out=./cli --go-grpc_out=./cli proto/*.proto
    @# printf "\033[0;33mCompiling proto files [Language: Rust] \033[0m \n"
    @# protoc --prost_out=client/src-tauri/src --tonic_out=client/src-tauri/src proto/*.proto

setup:
    @printf "\033[0;33mInstalling Golang Dependencies: Started \033[0m \n"
    @cd server && go mod tidy
    @printf "\033[0;33mInstalling Golang Dependencies: Done \033[0m \n"
    @printf "\033[0;33mGenerating the server's Private Key and Self-Assigned Certificate:Started \033[0m \n"
    @cd server && openssl genrsa -out server.pem 2048
    @cd server && openssl req -new -x509 -sha256 -key server.pem -out server.crt -days 3650
    @printf "\033[0;33mGenerating the server's Private Key and Self-Assigned Certificate:Done \033[0m \n"

run-server:
    @cd server && go run .

run-client-dev:
    @cd client && npm run tauri dev

run-client-build:
    @cd client && npm run tauri build