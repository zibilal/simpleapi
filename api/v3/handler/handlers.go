package handler

import (
	"github.com/zibilal/simpleapi/api"
	"github.com/zibilal/simpleapi/api/v3/response"
	"net/http"
)

func PingApi(engineContext api.EngineContext) error {
	// Unwrap the engineContext to the actual engineContext type

	val := engineContext.Get("MID")
	if val != nil {
		sval := val.(string)
		rsp := response.NewVersionOneBaseResponse(1000, "With Middleware "+sval)

		return engineContext.JSON(http.StatusOK, rsp)
	}

	rsp := response.NewVersionOneBaseResponse(1000, "Without Middleware")
	return engineContext.JSON(http.StatusOK, rsp)
}

func MiddlewareTest(engineContext api.EngineContext) error {
	engineContext.Set("MID", "here")
	return nil
}
