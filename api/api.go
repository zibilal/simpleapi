package api

type Engine interface {
	RegisterVersion(versions ...Version) error
	Execute(serve string) error
}

type Endpoint struct {
	Path    string
	Method  string
	Handler func(engineContext EngineContext)
}

type Version struct {
	endpoints []Endpoint
}

func NewVersion(endpoints []Endpoint) *Version {
	v := new(Version)
	v.endpoints = endpoints
	return v
}

func (v *Version) AddEndpoint(path, method string, handler func(engineContext EngineContext)) {
	v.endpoints = append(v.endpoints, Endpoint{
		Path:    path,
		Method:  method,
		Handler: handler,
	})
}

func (v *Version) Router() []Endpoint {
	return v.endpoints
}

type EngineContext interface {
	BindJSON(output interface{}) error
	BindQuery(output interface{}) error
	BindUri(output interface{}) error
	BindForm(output interface{}) error
	UnwrapContext() interface{}
}

type RouterEngine interface {
	RegisterVersion(...Version) error
	Execute(string) error
}