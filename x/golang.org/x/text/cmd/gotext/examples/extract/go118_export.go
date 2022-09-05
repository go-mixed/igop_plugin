// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18

package main

import (
	_ "golang.org/x/text/cmd/gotext/examples/extract"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "main",
		Path: "golang.org/x/text/cmd/gotext/examples/extract",
		Deps: map[string]string{
			"golang.org/x/text/language":        "language",
			"golang.org/x/text/message":         "message",
			"golang.org/x/text/message/catalog": "catalog",
		},
	})
}
