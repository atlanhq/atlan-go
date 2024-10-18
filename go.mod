module github.com/atlanhq/atlan-go

go 1.22.0

toolchain go1.22.2

require (
	github.com/k0kubun/go-ansi v0.0.0-20180517002512-3bf9e2903213
	github.com/matoous/go-nanoid v1.5.0
	github.com/schollz/progressbar/v3 v3.14.4
	github.com/stretchr/testify v1.9.0
	// NOTE: We need to pin this experimental version of "slog" since it is compatible with Go 1.19
	// This is required because atlan-heracles uses go-sdk, which currently supports Go 1.19
	golang.org/x/exp v0.0.0-20241004190924-225e2abe05e6
)

require gonum.org/v1/gonum v0.15.1

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mitchellh/colorstring v0.0.0-20190213212951-d06e56a500db // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/term v0.20.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
