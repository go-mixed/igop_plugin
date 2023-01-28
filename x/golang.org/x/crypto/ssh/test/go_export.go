// export by github.com/goplus/igop/cmd/qexp

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
