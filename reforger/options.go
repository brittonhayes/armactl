package reforger

// ServerOption configures a ReforgerServer
type ServerOption func(*ReforgerServer)

// WithExecutable sets the path to ArmaReforgerServer.exe
func WithExecutable(exe string) ServerOption {
	return func(s *ReforgerServer) {
		s.Executable = exe
	}
}

// WithConfig sets the path to the server config JSON file
func WithConfig(config string) ServerOption {
	return func(s *ReforgerServer) {
		s.Config = config
	}
}

// WithMaxFPS sets the maximum FPS for the server
func WithMaxFPS(fps int) ServerOption {
	return func(s *ReforgerServer) {
		s.MaxFPS = fps
	}
}

// WithProfile sets the profile directory path
func WithProfile(profile string) ServerOption {
	return func(s *ReforgerServer) {
		s.Profile = profile
	}
}

// WithDetach sets whether to run the server in detached mode
func WithDetach(detach bool) ServerOption {
	return func(s *ReforgerServer) {
		s.Detach = detach
	}
}

// WithDryRun sets whether to run in dry-run mode
func WithDryRun(dryRun bool) ServerOption {
	return func(s *ReforgerServer) {
		s.DryRun = dryRun
	}
}

// WithSkipValidation sets whether to skip config validation
func WithSkipValidation(skip bool) ServerOption {
	return func(s *ReforgerServer) {
		s.SkipValidation = skip
	}
}

// WithValidator sets a custom config validator
func WithValidator(v ConfigValidator) ServerOption {
	return func(s *ReforgerServer) {
		s.Validator = v
	}
}

// WithInstallDir sets the server installation directory for SteamCMD
func WithInstallDir(dir string) ServerOption {
	return func(s *ReforgerServer) {
		s.InstallDir = dir
	}
}

// WithSteamUser sets the Steam username for SteamCMD authentication
func WithSteamUser(user string) ServerOption {
	return func(s *ReforgerServer) {
		s.SteamUser = user
	}
}

// WithSteamPassword sets the Steam password for SteamCMD authentication
func WithSteamPassword(password string) ServerOption {
	return func(s *ReforgerServer) {
		s.SteamPassword = password
	}
}

// WithSteamBranch sets the Steam branch (e.g., "experimental")
func WithSteamBranch(branch string) ServerOption {
	return func(s *ReforgerServer) {
		s.SteamBranch = branch
	}
}

// WithSteamCmdPath sets the path to the steamcmd executable
func WithSteamCmdPath(path string) ServerOption {
	return func(s *ReforgerServer) {
		s.SteamCmdPath = path
	}
}
