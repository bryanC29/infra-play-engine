package types

type Design struct {
	Nodes       []Node       `json:"nodes"`       // All components in the architecture
	Connections []Connection `json:"connections"` // Directed edges between nodes
}

// Node defines a single infrastructure component in the design.
type Node struct {
	ID        string        `json:"id"`   // Unique identifier for the node
	Type      string        `json:"type"` // E.g., "Database", "App", "Cache"
	Resources NodeResources `json:"resources"`
}

// NodeResources represents the resource configuration for a component.
type NodeResources struct {
	CPU      float64 `json:"cpu"`      // Per replica (in vCPUs)
	MemoryMB int     `json:"memoryMB"` // Per replica (in MB)
	Replicas int     `json:"replicas"` // Number of instances of the node
}

// Connection represents a directed edge from one node to another.
type Connection struct {
	From string `json:"from"` // Source node ID
	To   string `json:"to"`   // Target node ID
}
