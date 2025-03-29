package response

import "github.com/lrbell17/astroapi/impl/conf"

// Generic response interface
type Response[T any] interface {
	ResponseFromDao(dao *T, datasourceConf *conf.Datasource)
}
