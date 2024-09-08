package core

type Bootstrapper interface {
	CollectAdditionalOptions()
	BuildModel() *Directory
	Bootstrap()
}

type Directory struct {
	Name        string
	Files       []File
	Directories []Directory
}

type File struct {
	Name    string
	Content string
}
