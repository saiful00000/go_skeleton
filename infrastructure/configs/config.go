package configs

var LocalConfig *Config

type Config struct {
	DbUser string `mapstructure:"DBUSER"`
	DbPass string `mapstructure:"DBPASS"`
	DbIp   string `mapstructure:"DBIP"`
	DbName string `mapstructure:"DBNAME"`
	Port   string `mapstructure:"PORT"`
}
