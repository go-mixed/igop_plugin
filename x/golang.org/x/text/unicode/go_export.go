// export by github.com/goplus/igop/cmd/qexp

package unicode

import (
	_ "golang.org/x/text/unicode"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "unicode",
		Path: "golang.org/x/text/unicode",
		Deps: map[string]string{},
	})
}
