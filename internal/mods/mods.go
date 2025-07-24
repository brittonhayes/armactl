// Package mods provides utilities for working with Arma 3 mods.
package mods

import (
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/maruel/natural"
)

func parseNameFromMeta(s string) string {
	p := regexp.MustCompile(`"([^"]*)"`)
	s = p.FindString(s)
	return strings.ReplaceAll(s, "\"", "")
}

// Meta returns the name of the mod from the meta.cpp file
func Meta(file string) (*string, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	match := parseNameFromMeta(string(b))

	return &match, nil
}

// FromDirectory returns a list of mods from a directory
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

		modName, err := Meta(file)
		if err != nil {
			return nil, err
		}

		mods = append(mods, *modName)
	}
	sort.Strings(natural.StringSlice(mods))

	return mods, nil

}
