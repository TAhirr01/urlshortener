package config

var Conf Config

type Config struct {
	Application Application `yaml:"application"`
}

type Application struct {
	MongoDB MongoDB `yaml:"mongo-db"`
}

type MongoDB struct {
	URI string `yaml:"mongo-uri"`
}
