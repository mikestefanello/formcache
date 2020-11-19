## Overview

Rough POC for form state management in Go.

This includes an example 3-step form that can store data between requests.

## Concept

When the initial form is submitted, the form data is set in the cache and assigned a new ID. That ID is then output in the form templates as a hidden value:

`<input type="hidden" name="token" value="{{.Content.Token}}"/>`

On all requests to this form, the payload is inspected to see if a token exists, and if so, attempts to load a cache entry with that ID.

## Form

The form is located in `handlers/steps.go`. `HTTPHandler.Steps()` is the entrypoint.

## Caching

The cache is in memory, and implemented in `cache/cache.go`.

## Demo

To demo, `go run main.go`, navigate to `http://localhost:9000/steps` in your browser.