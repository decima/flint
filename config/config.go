package config

type SecurityConfig struct {
	JWTSecret string
}
type Config struct {
	StoragePath string
	Security    SecurityConfig
}

// NewConfig creates a new Config instance. In a real application, this might
// load values from a file or environment variables. This function acts as a provider for fx.
func NewConfig() *Config {
	return &Config{
		StoragePath: "./data/",
		Security: SecurityConfig{
			JWTSecret: "supersecret",
		},
	}
}
