package main

// config.yaml structure
type ConfigYaml struct {
	Api []ApiStruct `yaml:"api"`
}

// api structure
type ApiStruct struct {
	Path string `yaml:"path"`
	File string `yaml:"file"`
}

// template map structure
type TemplateMap struct {
	Api []ApiStruct
}
