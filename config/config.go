package config

import "github.com/hashicorp/hcl/v2/hclsimple"

type Config struct {
	Server           Server           `hcl:"server,block"`
	Goprotoarchetype Goprotoarchetype `hcl:"goprotoarchetype,block"`
}

type Server struct {
	BindAddr string `hcl:"bind_addr,attr"`
}

// GetConfig loads config from an hcl file at the specified path
func GetConfig(path string) (Config, error) {
	config := Config{}
	err := hclsimple.DecodeFile(path, nil, &config)

	return config, err
}
