//go:build tools
// +build tools

//go:generate go run github.com/99designs/gqlgen generate

package tools

import _ "github.com/99designs/gqlgen"
