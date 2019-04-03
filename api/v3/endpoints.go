package v3

import (
	"github.com/zibilal/simpleapi/api"
	"github.com/zibilal/simpleapi/api/v3/handler"
	"net/http"
)

func VersionOneApi() *api.Version {
	endpoints := []api.Endpoint{
		{
			Path:    "/ping",
			Method:  http.MethodGet,
			Handler: handler.PingApi,
			Middlewares: []api.ApiHandlerFunc{
				handler.MiddlewareTest,
			},
		},
	}
	return api.NewVersion("v3", endpoints)
}
