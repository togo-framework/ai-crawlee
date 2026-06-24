<!-- togo-header -->
<div align="center">
  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />
  <h1>togo-framework/ai-crawlee</h1>
  <p>
    <a href="https://to-go.dev/marketplace"><img src="https://img.shields.io/badge/marketplace-to--go.dev-1FC7DC" alt="marketplace" /></a>
    <a href="https://pkg.go.dev/github.com/togo-framework/ai-crawlee"><img src="https://pkg.go.dev/badge/github.com/togo-framework/ai-crawlee.svg" alt="pkg.go.dev" /></a>
    <img src="https://img.shields.io/badge/license-MIT-blue" alt="MIT" />
  </p>
  <p><strong>Part of the <a href="https://to-go.dev">togo</a> framework.</strong></p>
</div>

## Install

```bash
togo install togo-framework/ai-crawlee
```

<!-- /togo-header -->

# ai-crawlee

A togo **AI data-source** plugin — crawl a website and extract page text for `ai-rag` ingest and agents. Crawlee is a Node library; this is a Go-native crawler (colly) serving the same purpose.

```
togo install togo-framework/ai-crawlee
```

## Use
- Go: `crawlee.FromKernel(k).Crawl(ctx, "https://example.com", 20)` → `[]Page{URL,Title,Text}`
- REST: `POST /api/ai/crawlee` `{"url":"…","maxPages":20}`

Part of the [togo AI kit](https://to-go.dev/ai). MIT.

<!-- togo-sponsors -->
---

<div align="center">
  <h3>Premium sponsors</h3>
  <p>
    <a href="https://id8media.com"><strong>ID8 Media</strong></a> &nbsp;·&nbsp;
    <a href="https://one-studio.co"><strong>One Studio</strong></a>
  </p>
  <p><sub>Support togo — <a href="https://github.com/sponsors/fadymondy">become a sponsor</a>.</sub></p>
</div>
<!-- /togo-sponsors -->
