build:
	GOOS=js GOARCH=wasm go build -o maps.wasm .

deploy: build
	cp index.html map.js maps.wasm wasm_exec.js ./public
	firebase deploy
