package util

import (
	"math"

	"github.com/baely/memap/internal/models"
)

func GetScaleXY(lat, lon float64, zoom float64) (float64, float64) {
	const a = 1.9743504858348
	const m = 50.0 / 27.0

	scaleX := m * math.Pow(a, zoom)
	scaleY := scaleX / math.Cos(lat*math.Pi/180.0)

	return scaleX, scaleY
}

func TranslateToPosition(lat, lon float64, zoom float64, width, height int, pos *models.Position) (int, int) {
	scaleX, scaleY := GetScaleXY(lat, lon, zoom)

	x := int(scaleX*(pos.Longitude-lon)) + width/2
	y := int(scaleY*(lat-pos.Latitude)) + height/2
	return x, y
}

func TranslateToLatLon(lat, lon float64, zoom float64, width, height int, x, y int) (float64, float64) {
	scaleX, scaleY := GetScaleXY(lat, lon, zoom)

	outLon := float64(x-width/2)/scaleX + lon
	outLat := float64(height/2-y)/scaleY + lat
	return outLat, outLon
}

func Distance(pos1, pos2 *models.Position) float64 {
	const R = 6371e3 // Earth radius in meters

	lat1 := pos1.Latitude * math.Pi / 180.0
	lat2 := pos2.Latitude * math.Pi / 180.0
	deltaLat := (pos2.Latitude - pos1.Latitude) * math.Pi / 180.0
	deltaLon := (pos2.Longitude - pos1.Longitude) * math.Pi / 180.0

	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(deltaLon/2)*math.Sin(deltaLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	d := R * c

	return d
}

// Distance2ToRay returns the squared distance from point (px, py) to the ray
// starting at (ox, oy) in direction (dx, dy). Only snaps forward (t > 0).
func Distance2ToRay(px, py, ox, oy, dx, dy int) (int, int, int) {
	dot := (px-ox)*dx + (py-oy)*dy
	lenSq := dx*dx + dy*dy

	if lenSq == 0 || dot <= 0 {
		return px*px + py*py, px, py // effectively infinite distance
	}

	t := float64(dot) / float64(lenSq)

	nearX := float64(ox) + t*float64(dx)
	nearY := float64(oy) + t*float64(dy)

	fx := float64(px) - nearX
	fy := float64(py) - nearY

	return int(fx*fx + fy*fy), int(nearX), int(nearY)
}
