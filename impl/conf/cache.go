package conf

type (
	Cache struct {
		Host        string           `yaml:"host"`
		Port        string           `yaml:"port"`
		Password    string           `yaml:"password"`
		Performance CachePerformance `yaml:"performance"`
	}
	CachePerformance struct {
		Expiry        int `yaml:"expiry"`
		MaxRetries    int `yaml:"max_retries"`
		RetryInterval int `yaml:"retry_interval"`
	}
)
