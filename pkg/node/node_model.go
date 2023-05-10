package node

type Node struct {
	Name            string `json:"name"`
	UUID            string `json:"uuid"`
	MemoryAllocated string `json:"memoryAllocated"`
	CPUAllocated    string `json:"cpuAllocated"`
}

type Metrics struct {
	Name        string `json:"name"`
	UUID        string `json:"uuid"`
	MemoryUsage string `json:"memoryUsage"`
	CPUUsage    string `json:"cpuUsage"`
}
