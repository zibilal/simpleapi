package wrapper

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zibilal/simpleapi/api"
	"net/http"
)

// GonicEngine is wrapper type for gin.Engine type
type GonicEngine struct {
	gonicEngine *gin.Engine
}

func NewGonicEngine() *GonicEngine {
	router := new(GonicEngine)
	router.gonicEngine = gin.Default()

	return router
}

func (e *GonicEngine) RegisterVersion(versionName string, versions ...api.Version) error {
	for _, version := range versions {
		routeVersion := e.gonicEngine.Group(versionName)
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

// GonicEngineContext is a wrapper type for gin.Context type
type GonicEngineContext struct {
	ctx *gin.Context
}

func WrapGinContext(ctx *gin.Context) *GonicEngineContext {
	gonicCtx := new(GonicEngineContext)

	return gonicCtx
}

func (c *GonicEngineContext) BindJSON(output interface{}) error {
	return c.ctx.BindJSON(output)
}

func (c *GonicEngineContext) BindQuery(output interface{}) error {
	return c.ctx.BindQuery(output)
}

func (c *GonicEngineContext) BindUri(output interface{}) error {
	return c.ctx.Bind(output)
}

func (c *GonicEngineContext) BindForm(output interface{}) error {
	return c.ctx.Bind(output)
}