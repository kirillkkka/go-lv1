package yamlconf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type YamlConfig struct {
	Port        string    `yaml:"port"`
	DBUrl       string `yaml:"db_url"`
	JaegerUrl   string `yaml:"jaeger_url"`
	SentryUrl   string `yaml:"sentry_url"`
	KafkaBroker string `yaml:"kafka_broke"`
	AppID   string `yaml:"some_app_id"`
	AppKey  string `yaml:"some_app_key"`
}

func Parse() YamlConfig {
	file, err := os.Open("conf.yaml")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	conf := YamlConfig{}

	err = yaml.NewDecoder(file).Decode(&conf)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("Config uploaded successfully!")
	return conf
}