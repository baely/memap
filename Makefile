build:
	GOOS=js GOARCH=wasm go build -ldflags "-X main.EditMode=false" -o maps.wasm .

build-edit:
	GOOS=js GOARCH=wasm go build -ldflags "-X main.EditMode=true" -o editor.wasm .

build-all: build build-edit

deploy: build-all
	cp index.html map.js maps.wasm wasm_exec.js ./public
	firebase deploy
