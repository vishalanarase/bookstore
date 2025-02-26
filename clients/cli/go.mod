module github.com/vishalanarase/bookstore/clients/cli

go 1.21

require (
	github.com/olekukonko/tablewriter v0.0.5
	github.com/spf13/cobra v1.8.1
	github.com/vishalanarase/bookstore/clients/openapi v0.0.0-00010101000000-000000000000
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)

replace github.com/vishalanarase/bookstore/clients/openapi => ../openapi
