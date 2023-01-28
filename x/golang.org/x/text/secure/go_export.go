// export by github.com/goplus/igop/cmd/qexp

package secure

import (
	_ "golang.org/x/text/secure"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "secure",
		Path: "golang.org/x/text/secure",
		Deps: map[string]string{},
	})
}
