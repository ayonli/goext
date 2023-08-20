module github.com/ayonli/goext/oop

go 1.21.0

replace github.com/ayonli/goext/stringx/mbstring => ../stringx/mbstring

replace github.com/ayonli/goext/slicex => ../slicex

replace github.com/ayonli/goext/mathx => ../mathx

require (
	github.com/ayonli/goext/slicex v0.0.0-00010101000000-000000000000
	github.com/ayonli/goext/stringx v0.0.0-00010101000000-000000000000
	github.com/ayonli/goext/stringx/mbstring v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.8.4
)

require (
	github.com/ayonli/goext/mathx v0.0.0-00010101000000-000000000000 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/ayonli/goext/stringx => ../stringx
