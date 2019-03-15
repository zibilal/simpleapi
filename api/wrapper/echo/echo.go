package echo

import (
	"github.com/labstack/echo"
	"github.com/zibilal/simpleapi/api"
)

// EchoEngine is wrapper type for echo.Echo type
type EchoEngine struct {
	echoEngine *echo.Echo
}

func NewEchoEngine() *EchoEngine {
	router := new(EchoEngine)
	router.echoEngine = echo.New()

	return router
}

func(e *EchoEngine) RegisterVersion(versionName string, versions ...api.Version) error {
	for _, version := range versions {
		rootVersion := e.echoEngine.Group(versionName)
	}
}
