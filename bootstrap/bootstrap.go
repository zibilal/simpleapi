package bootstrap

type Bootstrap interface {
	Init() error
	Run()
}
