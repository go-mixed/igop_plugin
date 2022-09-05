// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18

package test

import (
	_ "golang.org/x/crypto/ssh/test"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "test",
		Path: "golang.org/x/crypto/ssh/test",
		Deps: map[string]string{},
	})
}
