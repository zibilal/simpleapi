package v1

import (
	"github.com/zibilal/simpleapi/api"
)

func VersionOneApi() *api.Version {
	endpoints := []api.Endpoint {
		{
			Path: "",
			Method: "",
			Handler: func(engineContext interface{}) {

			},
		},
	}
	return api.NewVersion("v1", endpoints)
}