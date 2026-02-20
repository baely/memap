package models

type Position struct {
	Latitude  float64
	Longitude float64
}

type Node struct {
	Id          int
	Label       string
	Link        string
	Description string
	Position    *Position
}

type PathType string

const (
	PathTypeRoad    PathType = "PathTypeRoad"
	PathTypeRailway PathType = "PathTypeRailway"
)

type Path struct {
	Id       int
	Label    string
	Type     PathType
	Nodes    []*Node
	NodeRefs []int
}

type Map struct {
	Paths []*Path
	Nodes []*Node
}
