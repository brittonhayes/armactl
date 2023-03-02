package preset

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Parse(t *testing.T) {
	want := &Preset{
		Name: "AntistasiServer",
		Mods: []string{
			"3CB Factions",
			"ACE Compat - RHS AFRF",
			"ACE Compat - RHS USAF",
			"ACE Compat - RHS: GREF",
			"ACE Compat - RHS: SAF",
			"Antistasi Plus",
			"CBA_A3",
			"CUP Terrains - Core",
			"CUP Terrains - Maps 2.0",
			"D3S Cars pack",
			"RHSAFRF",
			"RHSGREF",
			"RHSPKL",
			"RHSSAF",
			"RHSUSAF",
			"ace",
		},
	}

	b := bytes.NewBufferString(PRESET_HTML)
	got, err := New().Parse(b)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, got, want)
}
