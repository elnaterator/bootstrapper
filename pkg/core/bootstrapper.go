package core

type Bootstrapper interface {
	CollectAdditionalOptions()
	BuildModel() *Directory
	Bootstrap() error
}
