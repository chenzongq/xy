package config

type Server struct {
	AutoCode AutoCode `mapstructure:"autocode" json:"autocode" yaml:"autocode"`
	Zap      Zap      `mapstructure:"zap" json:"zap" yaml:"zap"`
	Mysql    Mysql    `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	System   System   `mapstructure:"system" json:"system" yaml:"system"`
	Timer    Timer    `mapstructure:"timer" json:"timer" yaml:"timer"`
	Redis    Redis    `mapstructure:"redis" json:"redis" yaml:"redis"`
	Email    Email    `mapstructure:"email" json:"email" yaml:"email"`
}
