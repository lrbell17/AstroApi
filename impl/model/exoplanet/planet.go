package exoplanet

// const (
// 	NameIdx     = 1
// 	HostIdx     = 2
// 	DistanceIdx = 10
// 	RadiusIdx   = 11
// 	MassIdx     = 13
// )

type (
	Exoplanet struct {
		ID     uint `gorm:"primaryKey"`
		Name   string
		Host   string
		Mass   float32
		Radius float32
		Dist   float32
	}
)
