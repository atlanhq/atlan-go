module github.com/atlanhq/atlan-go

go 1.19

require (
	github.com/matoous/go-nanoid v1.5.0
	github.com/stretchr/testify v1.9.0
	// NOTE: We need to pin this experimental version of "slog" since it is compatible with Go 1.19
	// This is required because atlan-heracles uses go-sdk, which currently supports Go 1.19
	golang.org/x/exp v0.0.0-20240707233637-46b078467d37
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
