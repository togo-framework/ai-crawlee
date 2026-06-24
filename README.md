# ai-crawlee

A togo **AI data-source** plugin — crawl a website and extract page text for `ai-rag` ingest and agents. Crawlee is a Node library; this is a Go-native crawler (colly) serving the same purpose.

```
togo install togo-framework/ai-crawlee
```

## Use
- Go: `crawlee.FromKernel(k).Crawl(ctx, "https://example.com", 20)` → `[]Page{URL,Title,Text}`
- REST: `POST /api/ai/crawlee` `{"url":"…","maxPages":20}`

Part of the [togo AI kit](https://to-go.dev/ai). MIT.
