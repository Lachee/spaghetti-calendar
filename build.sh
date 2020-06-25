echo "Building WASM..."
cd ./src/
GOOS=js GOARCH=wasm go build -o ../resources/noodle.wasm .
echo "Done"
