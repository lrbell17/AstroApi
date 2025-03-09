package stars

const modelName = "star"

type (
	Star struct {
		ID     uint `gorm:"primaryKey"`
		Name   string
		Mass   float32
		Radius float32
		Temp   float32
	}
)

func (*Star) GetModelName() string {
	return modelName
}

func (*Star) ValidateColumns(header map[string]int) error {
	// TODO
	return nil
}
