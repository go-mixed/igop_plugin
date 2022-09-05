// export by github.com/goplus/igop/cmd/qexp

//go:build go1.19

package pkg

import (
	q "golang.org/x/text/cmd/gotext/examples/extract_http/pkg"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "pkg",
		Path: "golang.org/x/text/cmd/gotext/examples/extract_http/pkg",
		Deps: map[string]string{
			"golang.org/x/text/language": "language",
			"golang.org/x/text/message":  "message",
			"net/http":                   "http",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Generize": reflect.ValueOf(q.Generize),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
