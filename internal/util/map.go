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

func SegmentIntersect(x1, y1, x2, y2, x3, y3, x4, y4 int) (float64, float64, bool) {
	dx1 := x2 - x1
	dy1 := y2 - y1
	dx2 := x4 - x3
	dy2 := y4 - y3

	denom := dx1*dy2 - dy1*dx2
	if denom == 0 {
		return 0, 0, false
	}

	tNum := (x3-x1)*dy2 - (y3-y1)*dx2
	sNum := (x3-x1)*dy1 - (y3-y1)*dx1

	t := float64(tNum) / float64(denom)
	s := float64(sNum) / float64(denom)

	const eps = 0.01
	if t <= eps || t >= 1-eps || s <= eps || s >= 1-eps {
		return 0, 0, false
	}

	return t, s, true
}
