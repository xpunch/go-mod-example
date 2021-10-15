package example

type Service interface {
	Run() error
}

type Plugin interface {
	Load() error
}
