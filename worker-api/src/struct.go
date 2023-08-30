package main

// config.yaml structure
type ConfigYaml struct {
	Manager  ManagerStruct  `yaml:"manager"`
	Server   ServerStruct   `yaml:"server"`
	RabbitMQ RabbitMQStruct `yaml:"rabbitmq"`
}

type ManagerStruct struct {
	Host string `yaml:"host"`
	Port string `yaml:"port" default:"80"`
}

type ServerStruct struct {
	Host string `yaml:"host"`
}

type RabbitMQStruct struct {
	Host     string         `yaml:"host"`
	Port     string         `yaml:"port" default:"80"`
	QoS      QoSStruct      `yaml:"qos"`
	Queue    QueueStruct    `yaml:"queue"`
	Consumer ConsumerStruct `yaml:"consumer"`
}

type QoSStruct struct {
	PrefetchCount int  `yaml:"prefetchCount"`
	PrefetchSize  int  `yaml:"prefetchSize"`
	Global        bool `yaml:"global"`
}

type QueueStruct struct {
	Name string `yaml:"name"`
}

type ConsumerStruct struct {
	Name string `yaml:"name" default:""`
}
