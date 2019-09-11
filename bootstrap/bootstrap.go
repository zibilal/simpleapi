package bootstrap

import "github.com/zibilal/simpleapi/appctx"

type Bootstrap interface {
	Init() error
	Run() error
	ApplicationContext() *appctx.AppContext
}
