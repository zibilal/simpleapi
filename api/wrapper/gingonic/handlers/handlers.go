package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/zibilal/logwrapper"
	"github.com/zibilal/simpleapi/api"
)

func PingApi(engineContext api.EngineContext) {
	// Unwrap the engineContext to the actual engineContext type
	eCtx, found := engineContext.UnwrapContext().(*gin.Context)

	if !found {
		logwrapper.Fatal("Wrong handler implementation. Please check your api implementation")
		return
	}

	eCtx.JSON(200, gin.H {
		"code": 1000,
		"message": "success",
	})
}