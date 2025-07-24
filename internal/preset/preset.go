// Package preset provides a parser for Arma 3 preset files.
package preset

import (
	"io"
	"sort"

	"github.com/PuerkitoBio/goquery"
	"github.com/maruel/natural"
)

type Preset struct {
	Name string
	Mods []string
}

type Parser interface {
	Parse(io.Reader) (*Preset, error)
}

func New() Parser {
	return &Preset{}
}

func (p *Preset) Parse(r io.Reader) (*Preset, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	p.Name = doc.Find("h1 > strong").Text()
	doc.Find("[data-type='ModContainer']").Each(func(i int, s *goquery.Selection) {
		p.Mods = append(p.Mods, s.Find("[data-type='DisplayName']").Text())
	})

	sort.Sort(natural.StringSlice(p.Mods))

	return p, nil
}
