run:
	go run .

build_wasm:
	env GOOS=js GOARCH=wasm go build -o web/assets/snakegame.wasm .