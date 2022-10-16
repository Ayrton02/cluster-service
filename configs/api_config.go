package configs

type APIConfig struct {
	Port           string
	AllowedOrigins string
}

func NewAPIConfig(port, allowedOrigins string) *APIConfig {
	return &APIConfig{
		Port:           port,
		AllowedOrigins: allowedOrigins,
	}
}
