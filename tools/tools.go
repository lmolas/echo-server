// +build tools

package tools

// tool dependencies
import (
	_ "github.com/swaggo/swag"
)

// If you want to use tools, please run the following command:
//  go generate ./tools.go
//
//go:generate go build -v -o=./bin/swag github.com/swaggo/swag/cmd/swag
