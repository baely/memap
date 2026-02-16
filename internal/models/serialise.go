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
	var allNodes []nodeKey
	usedNames := make(map[string]int)

	addNode := func(n *Node) string {
		k := nodeKey{n.Label, n.Position.Latitude, n.Position.Longitude}
		if v, ok := nodeToVar[k]; ok {
			return v
		}
		name := labelToIdent(n.Label)
		if count, exists := usedNames[name]; exists {
			usedNames[name] = count + 1
			name = fmt.Sprintf("%s%d", name, count+1)
		}
		usedNames[name]++
		nodeToVar[k] = name
		allNodes = append(allNodes, k)
		return name
	}

	// Register standalone nodes first, then path nodes.
	for _, n := range m.Nodes {
		addNode(n)
	}
	for _, p := range m.Paths {
		for _, n := range p.Nodes {
			addNode(n)
		}
	}

	// Build path variable names.
	pathVarNames := make([]string, len(m.Paths))
	usedPathNames := make(map[string]int)
	for i, p := range m.Paths {
		name := labelToIdent(p.Label)
		if count, exists := usedPathNames[name]; exists {
			usedPathNames[name] = count + 1
			name = fmt.Sprintf("%s%d", name, count+1)
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
		fmt.Fprintf(&buf, "%s = &Node{Label: %q, Position: &Position{%v, %v}}\n",
			name, k.Label, k.Lat, k.Lng)
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
func labelToIdent(label string) string {
	// Replace common separators with spaces for splitting.
	r := strings.NewReplacer("&", " ", "-", " ", "_", " ")
	label = r.Replace(label)
	words := strings.Fields(label)

	var parts []string
	for _, w := range words {
		// Strip non-alphanumeric characters.
		cleaned := strings.Map(func(r rune) rune {
			if unicode.IsLetter(r) || unicode.IsDigit(r) {
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
		return "n"
	}

	// First word fully lowered, subsequent words title-cased.
	var b strings.Builder
	b.WriteString(strings.ToLower(parts[0]))
	for _, p := range parts[1:] {
		b.WriteString(strings.ToUpper(p[:1]) + strings.ToLower(p[1:]))
	}

	ident := b.String()

	// If it starts with a digit, prefix with 'n'.
	if len(ident) > 0 && unicode.IsDigit(rune(ident[0])) {
		ident = "n" + ident
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
