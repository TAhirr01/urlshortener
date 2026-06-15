package config

var Conf Config

type Config struct {
	Application Application `yaml:"application"`
}

type Application struct {
	
}
