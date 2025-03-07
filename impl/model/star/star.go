package stars

type (
	Star struct {
		ID     uint `gorm:"primaryKey"`
		Name   string
		Mass   float32
		Radius float32
		Temp   float32
	}
)
