package models

type Position struct {
	Latitude  float64
	Longitude float64
}

type Node struct {
	Label    string
	Position Position
}

type PathType int

const (
	PathTypeRoad PathType = iota
	PathTypeRailway
)

type Path struct {
	Type  PathType
	Nodes []Node
}

type Map struct {
	Paths []Path
	Nodes []Node
}
