package util

import (
	"syscall/js"

	"github.com/baely/memap/internal/models"
)

func Download(m *models.Map) {
	data, err := models.SerialiseMap(m, "models", "SampleMap")
	if err != nil {
		js.Global().Get("console").Call("log", err.Error())
	}

	filename := "map.go"

	// Create a Uint8Array from your byte slice
	uint8Array := js.Global().Get("Uint8Array").New(len(data))
	js.CopyBytesToJS(uint8Array, data)

	// Create a blob from the array
	array := js.Global().Get("Array").New(uint8Array)
	blob := js.Global().Get("Blob").New(array)

	// Create an object URL for the blob
	url := js.Global().Get("URL").Call("createObjectURL", blob)

	// Create a temporary anchor element and trigger download
	a := js.Global().Get("document").Call("createElement", "a")
	a.Set("href", url)
	a.Set("download", filename)
	a.Call("click")

	// Clean up the object URL
	js.Global().Get("URL").Call("revokeObjectURL", url)
}

func DownloadFn(m *models.Map) func() {
	return func() {
		Download(m)
	}
}
