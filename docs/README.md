# ai-crawlee — documentation

Go-native site crawler data-source for the togo AI kit

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

Environment variables read by this plugin (extracted from the source — see the gateway/provider docs for each value):

_No environment variables read directly (uses the kernel/base config or the app DB)._

## Usage

```go
// A data source for ai-rag / agents: fetch/scrape/search web content.
docs, err := crawlee.FromKernel(k).Fetch(ctx, "https://example.com")
```

## Links

- Marketplace: https://to-go.dev/marketplace
- Source: https://github.com/togo-framework/ai-crawlee
- Full README: ../README.md
