package reforger

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {

	t.Run("valid config", func(t *testing.T) {
		file, err := os.ReadFile("../test/reforger.server.json")
		if err != nil {
			t.Fatal(err)
		}

		var config interface{}
		err = json.Unmarshal(file, &config)
		if err != nil {
			t.Fatal(err)
		}

		validator := NewConfigValidator()
		err = validator.Validate(context.Background(), config)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("invalid config", func(t *testing.T) {
		file, err := os.ReadFile("../test/invalid-reforger.server.json")
		if err != nil {
			t.Fatal(err)
		}

		var config interface{}
		err = json.Unmarshal(file, &config)
		if err != nil {
			t.Fatal(err)
		}

		validator := NewConfigValidator()
		assert.Error(t, validator.Validate(context.Background(), config))
	})
}
