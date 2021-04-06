package config

type System struct {
	Env                string `mapstructure:"env" json:"env" yaml:"env"`
	Addr               int    `mapstructure:"addr" json:"addr" yaml:"addr"`
	DbType             string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`
	OssType            string `mapstructure:"oss-type" json:"ossType" yaml:"oss-type"`
	UseMultipoint      bool   `mapstructure:"use-multipoint" json:"useMultipoint" yaml:"use-multipoint"`
	AutoRegisterRouter bool   `mapstructure:"auto-register-router" json:"auto-register-router" yaml:"auto-register-router"`
}
