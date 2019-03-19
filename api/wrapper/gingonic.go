package wrapper

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo"
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
					_ = r.Handler(WrapGinContext(c))
				})
			case http.MethodGet:
				routeVersion.GET(r.Path, func(c *gin.Context){
					_ = r.Handler(WrapGinContext(c))
				})
			case http.MethodPut:
				routeVersion.PUT(r.Path, func(c *gin.Context) {
					_ = r.Handler(WrapGinContext(c))
				})
			case http.MethodDelete:
				routeVersion.DELETE(r.Path, func(c *gin.Context) {
					_ = r.Handler(WrapGinContext(c))
				})
			case http.MethodPatch:
				routeVersion.PATCH(r.Path, func(c *gin.Context) {
					_ = r.Handler(WrapGinContext(c))
				})
			default:
				return errors.New("invalid version " + versionName	 + " unknown method " + r.Method)
			}
		}
	}
	return nil
}

func (e *GonicEngine) wrapHandler(handler api.ApiHandlerFunc, middlewares ...api.ApiHandlerFunc) []gin.HandlerFunc {
	var result []gin.HandlerFunc

	if handler == nil {
		return nil
	}

	result = make([]gin.HandlerFunc, 0)

	for _, m := range middlewares {
		result = append(result, func(c *gin.Context) {
			err := m(WrapGinContext(c))
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, )
			}
		})
	}

	result = append(result, func(c *gin.Context) {
		_ = handler(WrapGinContext(c))
	})
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
	gonicCtx.ctx = ctx
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