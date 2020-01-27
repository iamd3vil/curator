package main

import (
	"log"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
)

// Config is the app config
type Config struct {
	InputDir    string `koanf:"input_dir"`
	OutputDir   string `koanf:"output_dir"`
	CalibrePath string `koanf:"calibre_path"`
}

var (
	k   = koanf.New(".")
	cfg Config
)

func initConfig() {
	if err := k.Load(file.Provider("curator.toml"), toml.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	k.Unmarshal("curator", &cfg)
}
