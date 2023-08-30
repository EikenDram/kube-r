package main

type ConfigYaml struct {
	Database     DatabaseStruct      `yaml:"database"`
	RabbitMQ     RabbitMQStruct      `yaml:"rabbitmq"`
	Applications []ApplicationStruct `yaml:"applications"`
}

type DatabaseStruct struct {
	Host string `yaml:"host"`
	Port string `yaml:"port" default:"80"`
	Name string `yaml:"name"`
}

type RabbitMQStruct struct {
	Host string `yaml:"host"`
	Port string `yaml:"port" default:"80"`
}

type ApplicationStruct struct {
	Name   string `yaml:"name"`
	URL    string `yaml:"url"`
	Notify string `yaml:"notify"`
	Result string `yaml:"result"`
	Prefix string `yaml:"prefix"`
}
