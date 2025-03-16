package response

import "github.com/lrbell17/astroapi/impl/conf"

// Generic response interface
type Response[T any] interface {
	ResponseFromModel(model *T, datasourceConf *conf.Datasource)
}
