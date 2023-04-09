package config

import "time"

// Config ...
type Config struct {
	Server   Server
	Weather  Weather
	Sentence Sentence
}

// Server ...
type Server struct {
	Phone string `yaml:"phone"`
	Title string `yaml:"title"`
	End   string `yaml:"end"`
	Form  string `yaml:"form"`
}

// Weather ...
type Weather struct {
	Host     string        `yaml:"host"`
	Token    string        `yaml:"token"`
	Version  string        `yaml:"version"`
	Location string        `yaml:"location"`
	Timeout  time.Duration `yaml:"timeout"`
}

// Sentence ...
type Sentence struct {
	ShanBayURL  string        `yaml:"shanbay_url" mapstructure:"shanbay"`
	ICiBaURL    string        `yaml:"iciba_url" mapstructure:"iciba"`
	ShaDiaoHost string        `yaml:"shadiao_host" mapstructure:"shadiao"`
	Timeout     time.Duration `yaml:"timeout"`
}
