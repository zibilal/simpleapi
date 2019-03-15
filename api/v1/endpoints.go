package v1

import (
	"github.com/zibilal/simpleapi/api"
	"github.com/zibilal/simpleapi/api/wrapper/gingonic/handlers"
	"net/http"
)

func VersionOneApi() *api.Version {
	endpoints := []api.Endpoint {
		{
			Path: "/ping",
			Method: http.MethodGet,
			Handler: handlers.PingApi,
		},
	}
	return api.NewVersion(endpoints)
}