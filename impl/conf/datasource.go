package conf

type (
	Datasource struct {
		File          string        `yaml:"file"`
		ExoplanetData ExoplanetData `yaml:"exoplanet_data"`
		StarData      StarData      `yaml:"star_data"`
	}
	ExoplanetData struct {
		NameCol   string `yaml:"name"`
		HostCol   string `yaml:"host"`
		MassCol   string `yaml:"mass"`
		RadiusCol string `yaml:"radius"`
		DistCol   string `yaml:"dist"`
	}
	StarData struct {
		NameCol   string `yaml:"name"`
		MassCol   string `yaml:"mass"`
		RadiusCol string `yaml:"radius"`
		TempCol   string `yaml:"temp"`
	}
)
