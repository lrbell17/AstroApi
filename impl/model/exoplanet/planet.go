package exoplanet

import (
	"fmt"

	"github.com/lrbell17/astroapi/impl/conf"
	log "github.com/sirupsen/logrus"
)

const modelName = "exoplanet"

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

func (*Exoplanet) GetModelName() string {
	return modelName
}

func (*Exoplanet) ValidateColumns(header map[string]int) error {

	conf, _ := conf.GetConfig()
	requiredCols := []string{
		conf.Datasource.ExoplanetData.NameCol,
		conf.Datasource.ExoplanetData.HostCol,
		conf.Datasource.ExoplanetData.MassCol,
		conf.Datasource.ExoplanetData.RadiusCol,
		conf.Datasource.ExoplanetData.DistCol,
	}
	for _, col := range requiredCols {
		if _, ok := header[col]; !ok {
			return fmt.Errorf("required column %v does not exist", col)
		}
	}
	log.Infof("Column validation successsful: %v", requiredCols)
	return nil
}
