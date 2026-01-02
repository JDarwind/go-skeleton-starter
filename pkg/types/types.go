package types

type Server struct {
	Port   string
	Prefix string
}

type ProjectConfig struct {
	Server Server
}

type Config struct {
	ProjectConfig      *ProjectConfig
	ApplicationConfigs any
}
