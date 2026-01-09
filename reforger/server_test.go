package reforger

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDiscoverExecutable(t *testing.T) {
	t.Run("explicit path exists", func(t *testing.T) {
		// Create temp file
		tmpFile, err := os.CreateTemp("", "test-server-*.exe")
		if err != nil {
			t.Fatalf("failed to create temp file: %v", err)
		}
		defer os.Remove(tmpFile.Name())
		tmpFile.Close()

		exe, err := DiscoverExecutable(tmpFile.Name())
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
		if !strings.Contains(exe, filepath.Base(tmpFile.Name())) {
			t.Errorf("expected path to contain %s, got: %s", filepath.Base(tmpFile.Name()), exe)
		}
	})

	t.Run("explicit path not found", func(t *testing.T) {
		_, err := DiscoverExecutable("/nonexistent/path/server.exe")
		if err == nil {
			t.Error("expected error for nonexistent path, got nil")
		}
		if !strings.Contains(err.Error(), "not found") {
			t.Errorf("expected 'not found' in error, got: %v", err)
		}
	})

	t.Run("env var fallback", func(t *testing.T) {
		tmpFile, err := os.CreateTemp("", "test-server-*.exe")
		if err != nil {
			t.Fatalf("failed to create temp file: %v", err)
		}
		defer os.Remove(tmpFile.Name())
		tmpFile.Close()

		os.Setenv("ARMA_REFORGER_SERVER_EXE", tmpFile.Name())
		defer os.Unsetenv("ARMA_REFORGER_SERVER_EXE")

		exe, err := DiscoverExecutable("")
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
		if !strings.Contains(exe, filepath.Base(tmpFile.Name())) {
			t.Errorf("expected path to contain %s, got: %s", filepath.Base(tmpFile.Name()), exe)
		}
	})

	t.Run("no executable found", func(t *testing.T) {
		// Unset env var to ensure clean test
		os.Unsetenv("ARMA_REFORGER_SERVER_EXE")

		_, err := DiscoverExecutable("")
		if err == nil {
			t.Error("expected error when no executable found, got nil")
		}
		if !strings.Contains(err.Error(), "not found") {
			t.Errorf("expected 'not found' in error, got: %v", err)
		}
	})
}

func TestBuildArgs(t *testing.T) {
	t.Run("minimal config", func(t *testing.T) {
		s := &ReforgerServer{
			Config: "test.json",
		}

		args := s.buildArgs()
		if !strings.Contains(args, "-config=") {
			t.Error("expected args to contain -config=")
		}
		if !strings.Contains(args, "test.json") {
			t.Error("expected args to contain test.json")
		}
	})

	t.Run("with all options", func(t *testing.T) {
		s := &ReforgerServer{
			Config:  "test.json",
			MaxFPS:  120,
			Profile: "C:\\Profiles\\Server",
		}

		args := s.buildArgs()

		// Check config is present
		if !strings.Contains(args, "-config=") {
			t.Error("expected args to contain -config=")
		}

		// Check other args present
		if !strings.Contains(args, "-maxFPS=120") {
			t.Error("expected args to contain -maxFPS=120")
		}
		if !strings.Contains(args, "-profile=") {
			t.Error("expected args to contain -profile=")
		}
	})

	t.Run("maxFPS zero value not included", func(t *testing.T) {
		s := &ReforgerServer{
			Config: "test.json",
			MaxFPS: 0,
		}

		args := s.buildArgs()
		if strings.Contains(args, "-maxFPS=") {
			t.Error("expected args to not contain -maxFPS= when MaxFPS is 0")
		}
	})
}

func TestValidateConfig(t *testing.T) {
	t.Run("valid config", func(t *testing.T) {
		s := &ReforgerServer{
			Config:    filepath.Join("..", "test", "reforger.server.json"),
			Validator: NewConfigValidator(),
		}

		err := s.validateConfig(context.Background())
		if err != nil {
			t.Errorf("expected valid config to pass, got error: %v", err)
		}
	})

	t.Run("invalid config", func(t *testing.T) {
		s := &ReforgerServer{
			Config:    filepath.Join("..", "test", "invalid-reforger.server.json"),
			Validator: NewConfigValidator(),
		}

		err := s.validateConfig(context.Background())
		if err == nil {
			t.Error("expected invalid config to fail validation")
		}
	})

	t.Run("config file not found", func(t *testing.T) {
		s := &ReforgerServer{
			Config:    "nonexistent.json",
			Validator: NewConfigValidator(),
		}

		err := s.validateConfig(context.Background())
		if err == nil {
			t.Error("expected error for nonexistent config file")
		}
	})
}

func TestFunctionalOptions(t *testing.T) {
	s := New(
		WithExecutable("test.exe"),
		WithConfig("config.json"),
		WithMaxFPS(144),
		WithProfile("profile"),
		WithDetach(true),
		WithDryRun(true),
	).(*ReforgerServer)

	if s.Executable != "test.exe" {
		t.Errorf("expected Executable='test.exe', got '%s'", s.Executable)
	}
	if s.Config != "config.json" {
		t.Errorf("expected Config='config.json', got '%s'", s.Config)
	}
	if s.MaxFPS != 144 {
		t.Errorf("expected MaxFPS=144, got %d", s.MaxFPS)
	}
	if s.Profile != "profile" {
		t.Errorf("expected Profile='profile', got '%s'", s.Profile)
	}
	if !s.Detach {
		t.Error("expected Detach=true")
	}
	if !s.DryRun {
		t.Error("expected DryRun=true")
	}
	if s.Validator == nil {
		t.Error("expected Validator to be initialized by default")
	}
}
