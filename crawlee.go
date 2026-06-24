// Package crawlee is a togo AI data-source plugin: crawl a website and extract
// page text so ai-rag ingest and agents can pull whole-site content. Crawlee is
// a Node library; this is a Go-native crawler (colly) serving the same purpose.
// Registers an "ai-crawlee" service + REST: POST /api/ai/crawlee {"url","maxPages"}.
package crawlee

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/gocolly/colly/v2"
	"github.com/togo-framework/togo"
)

// Page is one crawled page.
type Page struct {
	URL   string `json:"url"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

// Source crawls same-domain pages from a start URL.
type Source struct{}

// New returns a crawler Source.
func New() *Source { return &Source{} }

// Crawl visits start and same-domain links (depth 2) up to maxPages, returning text.
func (s *Source) Crawl(ctx context.Context, start string, maxPages int) ([]Page, error) {
	if maxPages <= 0 {
		maxPages = 20
	}
	var pages []Page
	c := colly.NewCollector(colly.MaxDepth(2))
	if u, err := url.Parse(start); err == nil {
		c.AllowedDomains = []string{u.Hostname()}
	}
	c.OnHTML("html", func(e *colly.HTMLElement) {
		if len(pages) >= maxPages {
			return
		}
		pages = append(pages, Page{
			URL:   e.Request.URL.String(),
			Title: e.ChildText("title"),
			Text:  strings.Join(strings.Fields(e.ChildText("body")), " "),
		})
	})
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if len(pages) < maxPages {
			_ = e.Request.Visit(e.Attr("href"))
		}
	})
	if err := c.Visit(start); err != nil {
		return nil, err
	}
	c.Wait()
	if len(pages) > maxPages {
		pages = pages[:maxPages]
	}
	return pages, nil
}

// FromKernel returns the registered Source, or nil.
func FromKernel(k *togo.Kernel) *Source {
	if v, ok := k.Get("ai-crawlee"); ok {
		if s, ok := v.(*Source); ok {
			return s
		}
	}
	return nil
}

func init() {
	togo.RegisterProviderFunc("ai-crawlee", togo.PriorityService, func(k *togo.Kernel) error {
		s := New()
		k.Set("ai-crawlee", s)
		mount(k.Router, s)
		return nil
	})
}

func mount(r chi.Router, s *Source) {
	r.Post("/api/ai/crawlee", func(w http.ResponseWriter, req *http.Request) {
		var b struct {
			URL      string `json:"url"`
			MaxPages int    `json:"maxPages"`
		}
		if err := json.NewDecoder(req.Body).Decode(&b); err != nil || b.URL == "" {
			http.Error(w, `{"error":"url required"}`, http.StatusBadRequest)
			return
		}
		pages, err := s.Crawl(req.Context(), b.URL, b.MaxPages)
		if err != nil {
			http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadGateway)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{"pages": pages})
	})
}
