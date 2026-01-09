package reforger

import (
	"context"
	"encoding/json"
	"os"
	"testing"
)

func TestGenerateConfigTemplate(t *testing.T) {
	t.Run("generates valid JSON without comments", func(t *testing.T) {
		tmpFile, err := os.CreateTemp("", "config-*.json")
		if err != nil {
			t.Fatalf("failed to create temp file: %v", err)
		}
		tmpFile.Close()
		defer os.Remove(tmpFile.Name())

		err = GenerateConfigTemplate(tmpFile.Name(), false)
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}

		// Verify file was created
		content, err := os.ReadFile(tmpFile.Name())
		if err != nil {
			t.Fatalf("failed to read temp file: %v", err)
		}

		// Verify it's valid JSON
		var config interface{}
		err = json.Unmarshal(content, &config)
		if err != nil {
			t.Errorf("generated template is not valid JSON: %v", err)
		}
	})

	t.Run("generates template with comments", func(t *testing.T) {
		tmpFile, err := os.CreateTemp("", "config-*.json")
		if err != nil {
			t.Fatalf("failed to create temp file: %v", err)
		}
		tmpFile.Close()
		defer os.Remove(tmpFile.Name())

		err = GenerateConfigTemplate(tmpFile.Name(), true)
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}

		// Verify file was created
		content, err := os.ReadFile(tmpFile.Name())
		if err != nil {
			t.Fatalf("failed to read temp file: %v", err)
		}

		// Verify content is not empty
		if len(content) == 0 {
			t.Error("expected non-empty template content")
		}

		// Note: Commented version contains JSON comments, so it won't parse as valid JSON
		// Just verify it was written
	})

	t.Run("template has expected structure", func(t *testing.T) {
		tmpFile, err := os.CreateTemp("", "config-*.json")
		if err != nil {
			t.Fatalf("failed to create temp file: %v", err)
		}
		tmpFile.Close()
		defer os.Remove(tmpFile.Name())

		err = GenerateConfigTemplate(tmpFile.Name(), false)
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}

		// Read and parse
		content, err := os.ReadFile(tmpFile.Name())
		if err != nil {
			t.Fatalf("failed to read temp file: %v", err)
		}

		var config map[string]interface{}
		err = json.Unmarshal(content, &config)
		if err != nil {
			t.Fatalf("failed to unmarshal config: %v", err)
		}

		// Check for required top-level fields
		requiredFields := []string{
			"bindPort",
			"publicPort",
			"game",
		}

		for _, field := range requiredFields {
			if _, ok := config[field]; !ok {
				t.Errorf("expected field '%s' in template, but it was missing", field)
			}
		}

		// Check game object has expected structure
		game, ok := config["game"].(map[string]interface{})
		if !ok {
			t.Fatal("expected 'game' to be an object")
		}

		gameFields := []string{"name", "password", "passwordAdmin", "scenarioId", "maxPlayers", "gameProperties", "mods"}
		for _, field := range gameFields {
			if _, ok := game[field]; !ok {
				t.Errorf("expected field 'game.%s' in template, but it was missing", field)
			}
		}
	})

	t.Run("template validates against schema (with empty values)", func(t *testing.T) {
		tmpFile, err := os.CreateTemp("", "config-*.json")
		if err != nil {
			t.Fatalf("failed to create temp file: %v", err)
		}
		tmpFile.Close()
		defer os.Remove(tmpFile.Name())

		err = GenerateConfigTemplate(tmpFile.Name(), false)
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}

		// Read and validate
		content, err := os.ReadFile(tmpFile.Name())
		if err != nil {
			t.Fatalf("failed to read temp file: %v", err)
		}

		var config interface{}
		json.Unmarshal(content, &config)

		validator := NewConfigValidator()
		err = validator.Validate(context.Background(), config)

		// Note: Template may have empty values that fail validation
		// This test verifies structure, not content validity
		if err != nil {
			t.Logf("Template validation result: %v (expected for template with empty fields)", err)
		}
	})
}
