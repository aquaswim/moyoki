package port

type Server interface {
	Start() error
	Stop() error
}
