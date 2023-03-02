package preset

import (
	"io"
	"os"
	"sort"

	"github.com/PuerkitoBio/goquery"
	"github.com/maruel/natural"
)

type Preset struct {
	Name string
	Mods []string
}

func FromFile(path string) (*Preset, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	preset, err := parse(file)
	if err != nil {
		return nil, err
	}

	return preset, nil
}

func parse(r io.Reader) (*Preset, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	preset := &Preset{}

	preset.Name = doc.Find("h1 > strong").Text()
	doc.Find("[data-type='ModContainer']").Each(func(i int, s *goquery.Selection) {
		preset.Mods = append(preset.Mods, s.Find("[data-type='DisplayName']").Text())
	})

	sort.Sort(natural.StringSlice(preset.Mods))

	return preset, nil
}
