package main

import (
	"flag"
	"io/ioutil"
	"net/http"

	"github.com/grepory/demo-helm-chart-app/internal"
	yaml "gopkg.in/yaml.v3"
)

var (
	cfgPath = flag.String("f", "./config.yaml", "Path to config file")
)

type appConfig struct {
	Message string `yaml:"message"`
}

func main() {
	flag.Parse()

	cfgBytes, err := ioutil.ReadFile(*cfgPath)
	if err != nil {
		panic(err)
	}

	cfg := appConfig{}

	if err := yaml.Unmarshal(cfgBytes, &cfg); err != nil {
		panic(err)
	}

	app := internal.App{Message: cfg.Message}

	http.ListenAndServe(":8080", app)
}
