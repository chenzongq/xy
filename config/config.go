package config

type Server struct {
	AutoCode AutoCode `mapstructure:"autocode" json:"autocode" yaml:"autocode"`
	Zap      Zap      `mapstructure:"zap" json:"zap" yaml:"zap"`
}
