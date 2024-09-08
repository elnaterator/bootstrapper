package core

type Bootstrapper interface {
	CollectAdditionalOptions()
	BuildModel() *Directory
	Bootstrap() error
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

func BuildFile(name string, content string) File {
	return File{
		Name:    name,
		Content: content,
	}
}
