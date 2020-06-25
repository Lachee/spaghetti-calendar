cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
go build
./webserver.exe -dir ../ -dir ../resources/shader/ -dir ../resources/ -cmd "update-content.bat" -filter **/*.go -resources ../resources/