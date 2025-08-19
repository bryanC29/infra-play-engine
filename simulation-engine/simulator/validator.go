package simulator

import (
	"simengine/types"
)

func FindIsolatedNodes(design types.Design) (bool) {
	connected := make(map[string]struct{})

	for _, conn := range design.Connections {
		connected[conn.From] = struct{}{}
		connected[conn.To] = struct{}{}
	}

	var isolated []string
	for _, node := range design.Nodes {
		if _, exists := connected[node.ID]; !exists {
			isolated = append(isolated, node.ID)
		}
	}

	return len(isolated) == 0
}