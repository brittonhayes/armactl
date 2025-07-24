package server

type ServerOption func(*ARMA)

func WithBinary(binary string) ServerOption {
	return func(s *ARMA) {
		s.Binary = binary
	}
}

func WithCDLC(cdlc string) ServerOption {
	return func(s *ARMA) {
		s.CDLC = cdlc
	}
}

func WithConfig(config string) ServerOption {
	return func(s *ARMA) {
		s.Config = config
	}
}

func WithPort(port uint16) ServerOption {
	return func(s *ARMA) {
		s.Port = port
	}
}

func WithSkipInstall(skipInstall bool) ServerOption {
	return func(s *ARMA) {
		s.SkipInstall = skipInstall
	}
}

func WithModsLocal(mods []string) ServerOption {
	return func(s *ARMA) {
		s.ModsLocal = mods
	}
}

func WithModsServer(mods []string) ServerOption {
	return func(s *ARMA) {
		s.ModsServer = mods
	}
}

func WithWorld(world string) ServerOption {
	return func(s *ARMA) {
		s.World = world
	}
}

func WithLimitFPS(fps uint8) ServerOption {
	return func(s *ARMA) {
		s.LimitFPS = fps
	}
}

func WithHeadlessClients(clients uint8) ServerOption {
	return func(s *ARMA) {
		s.HeadlessClients = clients
	}
}

func WithHeadlessClientsProfile(profile string) ServerOption {
	return func(s *ARMA) {
		s.HeadlessClientProfile = profile
	}
}

func WithHeadlessClientsPassword(password string) ServerOption {
	return func(s *ARMA) {
		s.HeadlessClientPassword = password
	}
}

func WithHeadlessClientsServer(server string) ServerOption {
	return func(s *ARMA) {
		s.HeadlessClientServer = server
	}
}

func WithParams(params string) ServerOption {
	return func(s *ARMA) {
		s.Params = params
	}
}

func WithProfile(profile string) ServerOption {
	return func(s *ARMA) {
		s.Profile = profile
	}
}

func WithSteamBranch(branch string) ServerOption {
	return func(s *ARMA) {
		s.SteamBranch = branch
	}
}

func WithSteamBranchPassword(password string) ServerOption {
	return func(s *ARMA) {
		s.SteamBranchPassword = password
	}
}

func WithSteamUser(user string) ServerOption {
	return func(s *ARMA) {
		s.SteamUser = user
	}
}

func WithSteamPassword(password string) ServerOption {
	return func(s *ARMA) {
		s.SteamPassword = password
	}
}

func WithDryRun(dryRun bool) ServerOption {
	return func(s *ARMA) {
		s.DryRun = dryRun
	}
}
