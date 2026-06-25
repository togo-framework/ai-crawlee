# ai-crawlee — documentation

  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />

## Overview

Package crawlee is a togo AI data-source plugin: crawl a website and extract
page text so ai-rag ingest and agents can pull whole-site content. Crawlee is
a Node library; this is a Go-native crawler (colly) serving the same purpose.
Registers an "ai-crawlee" service + REST: POST /api/ai/crawlee {"url","maxPages"}.

## Install

```bash
togo install togo-framework/ai-crawlee
```

A capability plugin — it self-registers on boot; no driver selector needed.

## Configuration

Environment variables read by this plugin (extracted from the source):

_No environment variables read directly (uses the kernel/base config or the app DB)._

## Usage

```go
// A data source for ai-rag / agents: fetch/scrape/search web content.
src := crawlee.FromKernel(k)
docs, err := src.Fetch(ctx, "https://example.com")
```

## Links

- Marketplace: https://to-go.dev/marketplace
- Source: https://github.com/togo-framework/ai-crawlee
- README: ../README.md
