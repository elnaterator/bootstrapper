package core

type Project struct {
	RootDir     string
	ProjectType string
}

type Bootstrapper interface {
	Bootstrap()
}
