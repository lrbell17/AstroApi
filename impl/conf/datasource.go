package conf

type (
	Datasource struct {
		File          string        `yaml:"file"`
		ExoplanetData ExoplanetData `yaml:"exoplanet_data"`
		StarData      StarData      `yaml:"star_data"`
	}
	ExoplanetData struct {
		Name   Column `yaml:"name"`
		Host   Column `yaml:"host"`
		Mass   Column `yaml:"mass"`
		Radius Column `yaml:"radius"`
		Dist   Column `yaml:"dist"`
	}
	StarData struct {
		Name   Column `yaml:"name"`
		Mass   Column `yaml:"mass"`
		Radius Column `yaml:"radius"`
		Temp   Column `yaml:"temp"`
	}
	Column struct {
		ColName string `yaml:"column"`
		Unit    string `yaml:"unit,omitempty"` // Omitempty in case a unit is not provided
	}
)
