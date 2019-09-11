package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zibilal/logwrapper"
	"github.com/zibilal/simpleapi/api"
)

func PingApi(engineContext api.EngineContext) error {
	// Unwrap the engineContext to the actual engineContext type
	eCtx, found := engineContext.UnwrapContext().(*gin.Context)

	if !found {
		logwrapper.Fatal("Wrong handler implementation. Please check your api implementation")
		return errors.New("wrong handler implementation. Please check your api implementation")
	}

	val, found := eCtx.Get("MID")

	if found {
		eCtx.JSON(200, gin.H{
			"code":    1000,
			"message": val.(string),
		})
		return nil
	}

	eCtx.JSON(200, gin.H{
		"code":    1000,
		"message": "success",
	})

	return nil
}

func MiddlewareTest(engineContext api.EngineContext) error {
	// Unwrap the engineContext to the actual engineContext type
	eCtx, found := engineContext.UnwrapContext().(*gin.Context)

	if !found {
		logwrapper.Fatal("Wrong handler implementation. Please check your api implementation")
		return errors.New("wrong handler implementation. Please check your api implementation")
	}

	eCtx.Set("MID", "here")

	return nil
}
