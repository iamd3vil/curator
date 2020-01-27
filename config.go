package main

import (
	"log"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
)

type cfgCurator struct {
	InputDir    string   `koanf:"input_dir"`
	OutputDir   string   `koanf:"output_dir"`
	CalibrePath string   `koanf:"calibre_path"`
	Emails      []string `koanf:"emails"`
}

type cfgMailer struct {
	ServerAddr   string `koanf:"server_addr"`
	ServerPort   int    `koanf:"server_port"`
	SMTPUsername string `koanf:"smtp_username"`
	SMTPPassword string `koanf:"smtp_password"`
	FromAddr     string `koanf:"from_addr"`
}

// Config is the app config
type Config struct {
	app    cfgCurator
	mailer cfgMailer
}

var (
	k   = koanf.New(".")
	cfg Config
)

func initConfig() {
	if err := k.Load(file.Provider("curator.toml"), toml.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	k.Unmarshal("curator", &cfg.app)
	k.Unmarshal("mailer", &cfg.mailer)
}
