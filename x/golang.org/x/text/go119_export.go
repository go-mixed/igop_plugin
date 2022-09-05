// export by github.com/goplus/igop/cmd/qexp

//go:build go1.19

package text

import (
	_ "golang.org/x/text"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "text",
		Path: "golang.org/x/text",
		Deps: map[string]string{},
	})
}
