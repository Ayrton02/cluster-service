package pod

type Pod struct {
	Name    string  `json:"name"`
	Status  string  `json:"status"`
	IP      string  `json:"ip"`
	UUID    string  `json:"uuid"`
	Metrics Metrics `json:"metrics"`
}

type Metrics struct {
	Memory string `json:"memory"`
	CPU    string `json:"cpu"`
}
