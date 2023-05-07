package node

type Node struct {
	Name   string `json:"name"`
	UUID   string `json:"uuid"`
	Memory string `json:"memory"`
	CPU    string `json:"cpu"`
}

type Metrics struct {
	Node
}
