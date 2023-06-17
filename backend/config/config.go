package config

import (
	el_test "github.com/anukul/xmon/backend/execution_client/test"
	"os"

	"gopkg.in/yaml.v3"
)

type Node struct {
	Name              string            `yaml:"name"`
	Kind              string            `yaml:"kind"`
	ExecutionLayerURL string            `yaml:"el_rpc_url"`
	ConsensusLayerURL string            `yaml:"cl_rpc_url"`
	AuthHeaders       map[string]string `yaml:"auth_headers"`
}

type Config struct {
	ServerAddress   string                   `yaml:"server_address"`
	DatabasePath    string                   `yaml:"database_path"`
	RefreshInterval uint                     `yaml:"refresh_interval"`
	ChainName       string                   `yaml:"chain_name"`
	Nodes           []Node                   `yaml:"nodes"`
	Testing         bool                     `yaml:"testing"`
	TestNodes       map[string]el_test.Chain `yaml:"test_nodes"`
}

func NewConfig() (*Config, error) {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
