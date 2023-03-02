package mods

import (
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/maruel/natural"
)

type mod struct{}

type Mod interface {
	Meta(file string) (*string, error)
}

func NewMod() Mod {
	return &mod{}
}

func parseNameFromMeta(s string) string {
	p := regexp.MustCompile(`"([^"]*)"`)
	s = p.FindString(s)
	return strings.ReplaceAll(s, "\"", "")
}

func (m *mod) Meta(file string) (*string, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	match := parseNameFromMeta(string(b))

	return &match, nil
}

func FromDirectory(path string) ([]string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	abs, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	mods := make([]string, 0)
	for _, f := range files {
		if !f.IsDir() {
			continue
		}

		file := filepath.Join(abs, f.Name(), "meta.cpp")
		if _, err := os.Stat(file); err != nil {
			continue
		}

		modName, err := NewMod().Meta(file)
		if err != nil {
			return nil, err
		}

		mods = append(mods, *modName)
	}
	sort.Strings(natural.StringSlice(mods))

	return mods, nil

}
