compile:
    @printf "\033[0;33mCompiling proto files [Language: Golang] \033[0m \n"
    @protoc --go_out=./server --go-grpc_out=./server proto/*.proto
    @printf "\033[0;33mCompiling proto files [Language: Rust] \033[0m \n"
    @protoc --prost_out=client/src-tauri/src --tonic_out=client/src-tauri/src proto/*.proto

setup:
    @printf "\033[0;33mInstalling NPM Dependencies: Started\033[0m \n"
    @cd client && npm i 
    @printf "\033[0;33mInstalling NPM Dependencies: Done\033[0m \n"
    @printf "\033[0;33mInstalling Cargo Dependencies: Started \033[0m \n"
    @cd client/src-tauri && cargo install --path .
    @printf "\033[0;33mInstalling Cargo Dependencies: Done \033[0m \n"
    @printf "\033[0;33mInstalling Golang Dependencies: Started \033[0m \n"
    @cd server && go mod tidy
    @printf "\033[0;33mInstalling Golang Dependencies: Done \033[0m \n"

run-server:
    @cd server && go run .

run-client-dev:
    @cd client && npm run tauri dev

run-client-build:
    @cd client && npm run tauri build