package config

type Server struct {
	JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Mysql   Mysql   `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	RunPort RunPort `mapstructure:"runport" json:"runport" yaml:"runport"`
}
