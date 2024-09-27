package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/mux"
	"github.com/shynn2/cmd-gram/internal/api"
)

var (
	configPath string
)

func init() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	path = fmt.Sprintf("%s\\pkg\\client\\postgresql\\api.toml", path[:len(path)-3])
	flag.StringVar(&configPath, "config-path", path, "path to config file")
}

func main() {
	flag.Parse()
	config := api.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	api := api.New(mux.NewRouter(), config)

	if err := api.Start(); err != nil {
		log.Fatal(err)
	}
}
