package conf

type (
	Database struct {
		Host        string      `yaml:"host"`
		Name        string      `yaml:"name"`
		Port        string      `yaml:"port"`
		User        string      `yaml:"user"`
		Pass        string      `yaml:"password"`
		Performance Performance `yaml:"performance"`
	}
	Performance struct {
		Batchsize int `yaml:"batchsize"`
	}
)
