package types

type Server struct {
	Prefix string `yaml:"prefix"`
}

type ProjectConfig struct {
	Server Server `yaml:"server"`
}