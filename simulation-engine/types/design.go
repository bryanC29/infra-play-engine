package types

type Design struct {
	Nodes       []Node       `json:"nodes" validate:"required,dive"`       // All components in the architecture
	Connections []Connection `json:"connections" validate:"required,dive"` // Directed edges between nodes
}

// Node defines a single infrastructure component in the design.
type Node struct {
	ID        string        `json:"id" validate:"required"`   // Unique identifier for the node
	Name      string        `json:"name" validate:"required"` // E.g., "Database", "App", "Cache"
	Resources NodeResources `json:"resources" validate:"required"`
}

// NodeResources represents the resource configuration for a component.
type NodeResources struct {
	CPU      float64 `json:"cpu" validate:"required"`      // Per replica (in vCPUs)
	MemoryMB int     `json:"memoryMB" validate:"required"` // Per replica (in MB)
	Replicas int     `json:"replicas" validate:"required"` // Number of instances of the node
}

// Connection represents a directed edge from one node to another.
type Connection struct {
	From string `json:"from" validate:"required"` // Source node ID
	To   string `json:"to" validate:"required"`   // Target node ID
}
