package core

type Project struct {
	RootDir     Directory
	RootDirName string
	Type        string
}

func (p *Project) SetRootDir(rootDir string) {
	p.RootDirName = rootDir
}

func (p *Project) SetType(projectType string) {
	p.Type = projectType
}
