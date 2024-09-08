package core

type Project struct {
	RootDir     Directory
	RootDirName string
	Type        string
	Language    string
	BuildSystem string
	Framework   string
}

func (p *Project) SetRootDir(rootDir string) {
	p.RootDirName = rootDir
}

func (p *Project) SetType(projectType string) {
	p.Type = projectType
}

func (p *Project) SetLanguage(language string) {
	p.Language = language
}

func (p *Project) SetBuildSystem(buildSystem string) {
	p.BuildSystem = buildSystem
}
