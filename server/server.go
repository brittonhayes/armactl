package server

import "context"

type server struct {
	binary                 string `env:"ARMA_BINARY"`
	port                   uint16 `env:"ARMA_PORT" default:"2302"`
	skipInstall            bool   `env:"SKIP_INSTALL" default:"false"`
	localMods              string `env:"MODS_LOCAL"`
	serverMods             string `env:"MODS_SERVER"`
	world                  string `env:"ARMA_WORLD"`
	limitFPS               uint8  `env:"ARMA_LIMITFPS" default:"60"` // 0 = unlimited
	headlessClients        uint8  `env:"HEADLESS_CLIENTS" default:"0"`
	headlessClientsProfile string `env:"HEADLESS_CLIENTS_PROFILE" default:"headless"`
	params                 string `env:"ARMA_PARAMS"`
	profile                string `env:"ARMA_PROFILE"`
	steamUser              string `env:"STEAM_USER"`
	steamPassword          string `env:"STEAM_PASSWORD"`
}

type Server interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

func NewServer() Server {
	return &server{}
}

func (s *server) Start(ctx context.Context) error {
	return nil
}

func (s *server) Stop(ctx context.Context) error {
	return nil
}
