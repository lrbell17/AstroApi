package conf

type (
	Api struct {
		JwkPath            string `yaml:"jwk_path"`
		CorsAllowedOrigins string `yaml:"cors_allowed_origins"`
	}
)
