//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=cfg.yaml openapi.yml

// Package app exists to anchor code generation at the repo root, so that
// cfg.yaml's relative output path resolves the same way for `go generate`
// and for manual runs from the repo root.
package app
