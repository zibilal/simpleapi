package wrapper

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zibilal/simpleapi/api"
	"net/http"
)

type GonicEngine struct {
	gonicEngine *gin.Engine
}

func NewGonicEngine() *GonicEngine {
	router := new(GonicEngine)
	router.gonicEngine = gin.Default()

	return router
}

func (e *GonicEngine) RegisterVersion(versions ...api.Version) error {
	for _, version := range versions {
		routeVersion := e.gonicEngine.Group(version.Name())
		for _, r := range version.Router() {
			switch r.Method {
			case http.MethodPost:
				routeVersion.POST(r.Path, func(c *gin.Context){
					r.Handler(c)
				})
			case http.MethodGet:
				routeVersion.GET(r.Path, func(c *gin.Context){
					r.Handler(c)
				})
			case http.MethodPut:
				routeVersion.PUT(r.Path, func(c *gin.Context) {
					r.Handler(c)
				})
			case http.MethodDelete:
				routeVersion.DELETE(r.Path, func(c *gin.Context) {
					r.Handler(c)
				})
			case http.MethodPatch:
				routeVersion.PATCH(r.Path, func(c *gin.Context) {
					r.Handler(c)
				})
			default:
				return errors.New("invalid version " + version.Name() + " unknown method " + r.Method)
			}
		}
	}
	return nil
}

func (e *GonicEngine) Execute(serve string) error {
	return e.gonicEngine.Run(serve)
}