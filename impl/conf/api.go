package conf

type (
	Api struct {
		JwtExpiry          int    `yaml:"jwt_expiry"`
		JwtDomain          string `yaml:"jwt_domain"`
		JwkPath            string `yaml:"jwk_path"`
		RSAPrivatePath     string `yaml:"rsa_private_path"`
		CorsAllowedOrigins string `yaml:"cors_allowed_origins"`
	}
)
