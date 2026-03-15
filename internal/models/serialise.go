package models

import (
	"fmt"
	"go/format"
	"strings"
	"unicode"
)

// SerialiseMap generates Go source code that declares a Map variable with all
// its nodes and paths as package-level variables. The output is gofmt'd.
func SerialiseMap(m *Map, packageName string, varName string) ([]byte, error) {
	// Collect all unique nodes (by label+position as key).
	type nodeKey struct {
		Label string
		Lat   float64
		Lng   float64
	}
	nodeToVar := make(map[nodeKey]string)
	nodeData := make(map[nodeKey]*Node)
	var allNodes []nodeKey
	usedNames := make(map[string]int)

	addNode := func(n *Node, prefix string) string {
		k := nodeKey{n.Label, n.Position.Latitude, n.Position.Longitude}
		if v, ok := nodeToVar[k]; ok {
			return v
		}
		base := labelToIdent(n.Label, prefix)
		name := base
		for usedNames[name] > 0 {
			usedNames[base]++
			name = fmt.Sprintf("%s_%d", base, usedNames[base])
		}
		usedNames[name]++
		nodeToVar[k] = name
		nodeData[k] = n
		allNodes = append(allNodes, k)
		return name
	}

	// Register standalone nodes first, then path nodes.
	poiSet := make(map[nodeKey]bool)
	for _, n := range m.Nodes {
		k := nodeKey{n.Label, n.Position.Latitude, n.Position.Longitude}
		poiSet[k] = true
		addNode(n, "n")
	}
	for _, p := range m.Paths {
		for _, n := range p.Nodes {
			k := nodeKey{n.Label, n.Position.Latitude, n.Position.Longitude}
			if poiSet[k] {
				continue
			}
			addNode(n, "p")
		}
	}

	// Build path variable names.
	pathVarNames := make([]string, len(m.Paths))
	usedPathNames := make(map[string]int)
	for i, p := range m.Paths {
		base := labelToIdent(p.Label, "p")
		name := base
		for usedPathNames[name] > 0 {
			usedPathNames[base]++
			name = fmt.Sprintf("%s_%d", base, usedPathNames[base])
		}
		usedPathNames[name]++
		pathVarNames[i] = name
	}

	var buf strings.Builder

	buf.WriteString("package " + packageName + "\n\n")

	// Emit node variables.
	buf.WriteString("var (\n")
	for _, k := range allNodes {
		name := nodeToVar[k]
		n := nodeData[k]
		fmt.Fprintf(&buf, "%s = &Node{Label: %q", name, k.Label)
		if n.Link != "" {
			fmt.Fprintf(&buf, ", Link: %q", n.Link)
		}
		if n.Description != "" {
			fmt.Fprintf(&buf, ", Description: %q", n.Description)
		}
		fmt.Fprintf(&buf, ", Position: &Position{%v, %v}}\n", k.Lat, k.Lng)
	}
	buf.WriteString(")\n\n")

	// Emit path variables.
	buf.WriteString("var (\n")
	for i, p := range m.Paths {
		fmt.Fprintf(&buf, "%s = &Path{Label: %q, Type: %s, Nodes: []*Node{\n",
			pathVarNames[i], p.Label, pathTypeExpr(p.Type))
		for _, n := range p.Nodes {
			k := nodeKey{n.Label, n.Position.Latitude, n.Position.Longitude}
			fmt.Fprintf(&buf, "%s,\n", nodeToVar[k])
		}
		buf.WriteString("}}\n")
	}
	buf.WriteString(")\n\n")

	// Emit the Map variable.
	fmt.Fprintf(&buf, "var %s = &Map{\n", varName)
	if len(m.Nodes) > 0 {
		buf.WriteString("Nodes: []*Node{\n")
		for _, n := range m.Nodes {
			k := nodeKey{n.Label, n.Position.Latitude, n.Position.Longitude}
			fmt.Fprintf(&buf, "%s,\n", nodeToVar[k])
		}
		buf.WriteString("},\n")
	}
	if len(m.Paths) > 0 {
		buf.WriteString("Paths: []*Path{\n")
		for _, name := range pathVarNames {
			fmt.Fprintf(&buf, "%s,\n", name)
		}
		buf.WriteString("},\n")
	}
	buf.WriteString("}\n")

	return format.Source([]byte(buf.String()))
}

// labelToIdent converts a human label like "Flinders & Spencer" into a
// camelCase Go identifier like "flindersSpencer".
func labelToIdent(label string, prefix string) string {
	// Replace common separators with spaces for splitting.
	r := strings.NewReplacer("&", " ", "-", " ", "_", " ")
	label = r.Replace(label)
	words := strings.Fields(label)

	var parts []string
	for _, w := range words {
		// Strip non-ASCII and non-alphanumeric characters.
		cleaned := strings.Map(func(r rune) rune {
			if r <= unicode.MaxASCII && (unicode.IsLetter(r) || unicode.IsDigit(r)) {
				return r
			}
			return -1
		}, w)
		if cleaned == "" {
			continue
		}
		parts = append(parts, cleaned)
	}

	if len(parts) == 0 {
		return strings.ToLower(prefix)
	}

	// First word fully lowered, subsequent words title-cased.
	var b strings.Builder
	b.WriteString(strings.ToLower(parts[0]))
	for _, p := range parts[1:] {
		runes := []rune(p)
		b.WriteString(strings.ToUpper(string(runes[:1])) + strings.ToLower(string(runes[1:])))
	}

	ident := b.String()

	// If it starts with a digit, add the prefix.
	if len(ident) > 0 && unicode.IsDigit(rune(ident[0])) {
		ident = strings.ToLower(prefix) + ident
	}

	return ident
}

// pathTypeExpr returns the Go expression string for a PathType constant.
func pathTypeExpr(pt PathType) string {
	switch pt {
	case PathTypeRoad:
		return "PathTypeRoad"
	case PathTypeRailway:
		return "PathTypeRailway"
	default:
		return fmt.Sprintf("PathType(%q)", string(pt))
	}
}
