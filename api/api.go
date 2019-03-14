package api

type Engine interface {
	RegisterVersion(versions ...Version) error
	Execute(serve string) error
}

type Endpoint struct {
	Path    string
	Method  string
	Handler func(engineContext interface{})
}

type Version struct {
	name      string
	endpoints []Endpoint
}

func NewVersion(name string, endpoints []Endpoint) *Version {
	v := new(Version)
	v.name = name
	v.endpoints = endpoints
	return v
}

func (v *Version) AddEndpoint(path, method string, handler func(engineContext interface{})) {
	v.endpoints = append(v.endpoints, Endpoint{
		Path:    path,
		Method:  method,
		Handler: handler,
	})
}

func (v *Version) Name() string {
	return v.name
}

func (v *Version) Router() []Endpoint {
	return v.endpoints
}

type EngineContext interface {
	BindJSON(output interface{}) error
	BindQuery(output interface{}) error
	BindUri(output interface{}) error
	BindForm(output interface{}) error
}