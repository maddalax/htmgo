package sitemap

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/go-chi/chi/v5"
	"strings"
)

type URL struct {
	Loc        string  `xml:"loc"`
	ChangeFreq string  `xml:"changefreq,omitempty"`
	Priority   float32 `xml:"priority,omitempty"`
}

type URLSet struct {
	XMLName xml.Name `xml:"urlset"`
	XmlNS   string   `xml:"xmlns,attr"`
	URLs    []URL    `xml:"url"`
}

func NewSitemap(urls []URL) *URLSet {
	return &URLSet{
		XmlNS: "https://www.sitemaps.org/schemas/sitemap/0.9",
		URLs:  urls,
	}
}

func serialize(sitemap *URLSet) ([]byte, error) {
	buffer := bytes.Buffer{}
	enc := xml.NewEncoder(&buffer)
	enc.Indent("", "  ")
	if err := enc.Encode(sitemap); err != nil {
		return make([]byte, 0), fmt.Errorf("could not encode sitemap: %w", err)
	}
	return buffer.Bytes(), nil
}

func Generate(router *chi.Mux) ([]byte, error) {
	routes := router.Routes()
	urls := []URL{
		{
			Loc:        "/",
			Priority:   0.5,
			ChangeFreq: "weekly",
		},
		{
			Loc:        "/docs",
			Priority:   1.0,
			ChangeFreq: "daily",
		},
		{
			Loc:        "/examples",
			Priority:   0.7,
			ChangeFreq: "daily",
		},
		{
			Loc:        "/html-to-go",
			Priority:   0.5,
			ChangeFreq: "weekly",
		},
	}

	for _, route := range routes {
		if strings.HasPrefix(route.Pattern, "/docs/") {
			urls = append(urls, URL{
				Loc:        route.Pattern,
				Priority:   1.0,
				ChangeFreq: "weekly",
			})
		}

		if strings.HasPrefix(route.Pattern, "/examples/") {
			urls = append(urls, URL{
				Loc:        route.Pattern,
				Priority:   0.7,
				ChangeFreq: "weekly",
			})
		}
	}

	for i, url := range urls {
		urls[i].Loc = fmt.Sprintf("%s%s", "https://htmgo.dev", url.Loc)
	}

	sitemap := NewSitemap(urls)
	return serialize(sitemap)
}
