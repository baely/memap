package interactor

import (
	"sort"
	"syscall/js"

	"github.com/baely/memap/internal/canvas"
	"github.com/baely/memap/internal/models"
	"github.com/baely/memap/internal/util"
)

type drawPath struct {
	// Canvas fields
	canvas js.Value

	// Renderer fields
	*canvas.Renderer

	// Panning fields
	holding            bool
	panning            bool
	startLat, startLon float64
	startX, startY     int

	// Draw path fields
	currentPath      *models.Path
	cursorX, cursorY int
}

func NewDrawPath(renderer *canvas.Renderer, canvas js.Value) Interactor {
	return &drawPath{
		canvas:   canvas,
		Renderer: renderer,
	}
}

func (d *drawPath) Init() interface{} {
	d.canvas.Get("style").Set("cursor", "crosshair")
	d.currentPath = nil
	return d
}

func (d *drawPath) GetMenuItems() []InteractorMenu {
	return editorMenuItems(d.CurrentMap)
}

func (d *drawPath) updatePathCallback(this js.Value, args []js.Value) interface{} {
	label := args[0].String()

	path := d.GetSelectedPath()

	if path == nil {
		return nil
	}

	path.Label = label

	d.Draw()

	return nil
}

func (d *drawPath) findSnap(x, y int, threshold2 int) (*models.Node, int, int, bool) {
	minDistance := d.Width*d.Width + d.Height*d.Height

	var snapNode *models.Node
	var snapX, snapY int
	found := false

	// Check all nodes that belong to paths
	for _, path := range d.CurrentMap.Paths {
		for _, node := range path.Nodes {
			nodeX, nodeY := util.TranslateToPosition(d.Lat, d.Lon, d.Zoom, d.Width, d.Height, node.Position)

			dx := x - nodeX
			dy := y - nodeY
			distance2 := dx*dx + dy*dy

			if distance2 <= threshold2 && distance2 < minDistance {
				minDistance = distance2
				snapNode = node
				snapX = nodeX
				snapY = nodeY
				found = true
			}
		}
	}

	if found {
		return snapNode, snapX, snapY, true
	}

	var bisectPath *models.Path
	var bisectIndex int

	for _, path := range d.CurrentMap.Paths {
		for i, node := range path.Nodes {
			if i == len(path.Nodes)-1 {
				break
			}

			nextNode := path.Nodes[i+1]

			x1, y1 := util.TranslateToPosition(d.Lat, d.Lon, d.Zoom, d.Width, d.Height, node.Position)
			x2, y2 := util.TranslateToPosition(d.Lat, d.Lon, d.Zoom, d.Width, d.Height, nextNode.Position)

			distance2, nearX, nearY := distance2ToLine(x, y, x1, y1, x2, y2)

			if distance2 <= threshold2 && distance2 < minDistance {
				minDistance = distance2
				snapX = nearX
				snapY = nearY
				bisectPath = path
				bisectIndex = i
				found = true
			}
		}
	}

	if found && bisectPath != nil {
		lat, lon := util.TranslateToLatLon(d.Lat, d.Lon, d.Zoom, d.Width, d.Height, snapX, snapY)
		newNode := &models.Node{
			Position: &models.Position{
				Latitude:  lat,
				Longitude: lon,
			},
		}
		nodes := bisectPath.Nodes
		newNodes := make([]*models.Node, 0, len(nodes)+1)
		newNodes = append(newNodes, nodes[:bisectIndex+1]...)
		newNodes = append(newNodes, newNode)
		newNodes = append(newNodes, nodes[bisectIndex+1:]...)
		bisectPath.Nodes = newNodes

		snapNode = newNode
		return snapNode, snapX, snapY, true
	}

	return nil, x, y, false
}

func (d *drawPath) findSnapPreview(x, y int, threshold2 int) (int, int) {
	minDistance := d.Width*d.Width + d.Height*d.Height
	snapX, snapY := x, y

	// Check path nodes
	for _, path := range d.CurrentMap.Paths {
		for _, node := range path.Nodes {
			nodeX, nodeY := util.TranslateToPosition(d.Lat, d.Lon, d.Zoom, d.Width, d.Height, node.Position)

			dx := x - nodeX
			dy := y - nodeY
			distance2 := dx*dx + dy*dy

			if distance2 <= threshold2 && distance2 < minDistance {
				minDistance = distance2
				snapX = nodeX
				snapY = nodeY
			}
		}
	}

	if snapX != x || snapY != y {
		return snapX, snapY
	}

	// Check trajectory rays from the last node of the current path
	if d.currentPath != nil && len(d.currentPath.Nodes) > 0 {
		lastNode := d.currentPath.Nodes[len(d.currentPath.Nodes)-1]
		lx, ly := util.TranslateToPosition(d.Lat, d.Lon, d.Zoom, d.Width, d.Height, lastNode.Position)

		for _, path := range d.CurrentMap.Paths {
			for i, node := range path.Nodes {
				if node != lastNode {
					continue
				}

				neighbors := make([]int, 0, 2)
				if i > 0 {
					neighbors = append(neighbors, i-1)
				}
				if i < len(path.Nodes)-1 {
					neighbors = append(neighbors, i+1)
				}

				for _, ni := range neighbors {
					neighbor := path.Nodes[ni]
					nx, ny := util.TranslateToPosition(d.Lat, d.Lon, d.Zoom, d.Width, d.Height, neighbor.Position)

					dirX := lx - nx
					dirY := ly - ny

					dirs := [][2]int{
						{dirX, dirY},
						{-dirY, dirX},
						{dirY, -dirX},
					}

					for _, dir := range dirs {
						dist2, nearX, nearY := util.Distance2ToRay(x, y, lx, ly, dir[0], dir[1])

						if dist2 <= threshold2 && dist2 < minDistance {
							minDistance = dist2
							snapX = nearX
							snapY = nearY
						}
					}
				}
			}
		}

		if snapX != x || snapY != y {
			return snapX, snapY
		}
	}

	// Check path segments
	for _, path := range d.CurrentMap.Paths {
		for i, node := range path.Nodes {
			if i == len(path.Nodes)-1 {
				break
			}

			nextNode := path.Nodes[i+1]

			x1, y1 := util.TranslateToPosition(d.Lat, d.Lon, d.Zoom, d.Width, d.Height, node.Position)
			x2, y2 := util.TranslateToPosition(d.Lat, d.Lon, d.Zoom, d.Width, d.Height, nextNode.Position)

			distance2, nearX, nearY := distance2ToLine(x, y, x1, y1, x2, y2)

			if distance2 <= threshold2 && distance2 < minDistance {
				minDistance = distance2
				snapX = nearX
				snapY = nearY
			}
		}
	}

	return snapX, snapY
}

type intersectionHit struct {
	t    float64
	node *models.Node
	path *models.Path
	idx  int
}

func (d *drawPath) checkIntersections() {
	if d.currentPath == nil || len(d.currentPath.Nodes) < 2 {
		return
	}

	lastIdx := len(d.currentPath.Nodes) - 1
	segStart := d.currentPath.Nodes[lastIdx-1]
	segEnd := d.currentPath.Nodes[lastIdx]

	sx1, sy1 := util.TranslateToPosition(d.Lat, d.Lon, d.Zoom, d.Width, d.Height, segStart.Position)
	sx2, sy2 := util.TranslateToPosition(d.Lat, d.Lon, d.Zoom, d.Width, d.Height, segEnd.Position)

	var hits []intersectionHit

	for _, path := range d.CurrentMap.Paths {
		if path == d.currentPath {
			continue
		}
		for i := 0; i < len(path.Nodes)-1; i++ {
			n1 := path.Nodes[i]
			n2 := path.Nodes[i+1]

			px1, py1 := util.TranslateToPosition(d.Lat, d.Lon, d.Zoom, d.Width, d.Height, n1.Position)
			px2, py2 := util.TranslateToPosition(d.Lat, d.Lon, d.Zoom, d.Width, d.Height, n2.Position)

			t, _, ok := util.SegmentIntersect(sx1, sy1, sx2, sy2, px1, py1, px2, py2)
			if ok {
				// Interpolate in lat/lon space
				lat := segStart.Position.Latitude + t*(segEnd.Position.Latitude-segStart.Position.Latitude)
				lon := segStart.Position.Longitude + t*(segEnd.Position.Longitude-segStart.Position.Longitude)
				node := &models.Node{
					Position: &models.Position{
						Latitude:  lat,
						Longitude: lon,
					},
				}
				hits = append(hits, intersectionHit{t, node, path, i})
			}
		}
	}

	if len(hits) == 0 {
		return
	}

	// Sort by t (distance along the new segment)
	sort.Slice(hits, func(i, j int) bool {
		return hits[i].t < hits[j].t
	})

	newMiddle := make([]*models.Node, 0, len(hits))
	for _, h := range hits {
		newMiddle = append(newMiddle, h.node)
	}
	head := make([]*models.Node, lastIdx)
	copy(head, d.currentPath.Nodes[:lastIdx])
	d.currentPath.Nodes = append(append(head, newMiddle...), segEnd)

	pathHits := map[*models.Path][]intersectionHit{}
	for _, h := range hits {
		pathHits[h.path] = append(pathHits[h.path], h)
	}
	for path, pHits := range pathHits {
		sort.Slice(pHits, func(i, j int) bool {
			return pHits[i].idx < pHits[j].idx
		})
		for offset, h := range pHits {
			insertIdx := h.idx + 1 + offset
			nodes := path.Nodes
			newNodes := make([]*models.Node, 0, len(nodes)+1)
			newNodes = append(newNodes, nodes[:insertIdx]...)
			newNodes = append(newNodes, h.node)
			newNodes = append(newNodes, nodes[insertIdx:]...)
			path.Nodes = newNodes
		}
	}
}

func (d *drawPath) MouseDown(this js.Value, args []js.Value) interface{} {
	event := args[0]
	x := event.Get("clientX").Int()
	y := event.Get("clientY").Int()

	d.holding = true
	d.startX = x
	d.startY = y
	d.startLat = d.Lat
	d.startLon = d.Lon

	return nil
}

func (d *drawPath) MouseMove(this js.Value, args []js.Value) interface{} {
	event := args[0]
	x := event.Get("clientX").Int()
	y := event.Get("clientY").Int()

	d.cursorX = x
	d.cursorY = y

	snapX, snapY := d.findSnapPreview(x, y, 625)

	if !d.holding {
		d.Draw()
		d.drawPreviewLine(snapX, snapY)
		d.DrawCursor(snapX, snapY, "blue")
		return nil
	}

	deltaX := x - d.startX
	deltaY := y - d.startY

	if !d.panning {
		if deltaX*deltaX+deltaY*deltaY > 25 {
			d.panning = true
		}
	}

	if !d.panning {
		return nil
	}

	scaleX, scaleY := util.GetScaleXY(d.Lat, d.Lon, d.Zoom)
	deltaLat := float64(-deltaY) / scaleY
	deltaLon := float64(deltaX) / scaleX

	d.Lat = d.startLat - deltaLat
	d.Lon = d.startLon - deltaLon

	d.Draw()
	d.drawPreviewLine(snapX, snapY)
	d.DrawCursor(snapX, snapY, "blue")

	return nil
}

func (d *drawPath) drawPreviewLine(x, y int) {
	if d.currentPath == nil || len(d.currentPath.Nodes) == 0 {
		return
	}

	lastNode := d.currentPath.Nodes[len(d.currentPath.Nodes)-1]
	lastX, lastY := util.TranslateToPosition(d.Lat, d.Lon, d.Zoom, d.Width, d.Height, lastNode.Position)

	d.DrawLine(lastX, lastY, x, y, 4, "blue")
	d.SendBatch()
}

func (d *drawPath) finishPath() {
	if d.currentPath != nil {
		if len(d.currentPath.Nodes) < 2 {
			// Remove path with fewer than 2 nodes
			for i, p := range d.CurrentMap.Paths {
				if p == d.currentPath {
					d.CurrentMap.Paths = append(d.CurrentMap.Paths[:i], d.CurrentMap.Paths[i+1:]...)
					break
				}
			}
		} else {
			d.SetSelectedPath(d.currentPath, true, d.updatePathCallback)
		}
		d.currentPath = nil
	}
}

func (d *drawPath) MouseUp(this js.Value, args []js.Value) interface{} {
	event := args[0]
	x := event.Get("clientX").Int()
	y := event.Get("clientY").Int()

	if !d.panning {
		if d.currentPath == nil {
			d.currentPath = &models.Path{
				Label: "",
				Type:  models.PathTypeRoad,
			}
			d.CurrentMap.Paths = append(d.CurrentMap.Paths, d.currentPath)
		}

		node, _, _, found := d.findSnap(x, y, 625)

		if found && node != nil {
			if len(d.currentPath.Nodes) > 0 && d.currentPath.Nodes[len(d.currentPath.Nodes)-1] == node {
				d.finishPath()
			} else {
				d.currentPath.Nodes = append(d.currentPath.Nodes, node)
				d.checkIntersections()
			}
		} else {
			lat, lon := util.TranslateToLatLon(d.Lat, d.Lon, d.Zoom, d.Width, d.Height, x, y)
			newNode := &models.Node{
				Position: &models.Position{
					Latitude:  lat,
					Longitude: lon,
				},
			}
			d.currentPath.Nodes = append(d.currentPath.Nodes, newNode)
			d.checkIntersections()
		}
	}

	d.holding = false
	d.panning = false

	d.Draw()

	return nil
}

func (d *drawPath) MouseLeave(this js.Value, args []js.Value) interface{} {
	d.panning = false
	d.holding = false

	return nil
}

func (d *drawPath) Wheel(this js.Value, args []js.Value) interface{} {
	event := args[0]
	deltaY := event.Get("deltaY").Int()

	const zoomSensitivity = 0.001

	d.Zoom -= float64(deltaY) * zoomSensitivity

	const minZoom = 14.0
	const maxZoom = 18.0

	if d.Zoom < minZoom {
		d.Zoom = minZoom
	}

	if d.Zoom > maxZoom {
		d.Zoom = maxZoom
	}

	d.Draw()
	d.DrawCursor(d.cursorX, d.cursorY, "blue")

	return nil
}
