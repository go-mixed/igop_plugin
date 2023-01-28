// export by github.com/goplus/igop/cmd/qexp

package context

import (
	q "golang.org/x/net/context"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "context",
		Path: "golang.org/x/net/context",
		Deps: map[string]string{
			"context": "context",
			"time":    "time",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{
			"CancelFunc": reflect.TypeOf((*q.CancelFunc)(nil)).Elem(),
			"Context":    reflect.TypeOf((*q.Context)(nil)).Elem(),
		},
		Vars: map[string]reflect.Value{
			"Canceled":         reflect.ValueOf(&q.Canceled),
			"DeadlineExceeded": reflect.ValueOf(&q.DeadlineExceeded),
		},
		Funcs: map[string]reflect.Value{
			"Background":   reflect.ValueOf(q.Background),
			"TODO":         reflect.ValueOf(q.TODO),
			"WithCancel":   reflect.ValueOf(q.WithCancel),
			"WithDeadline": reflect.ValueOf(q.WithDeadline),
			"WithTimeout":  reflect.ValueOf(q.WithTimeout),
			"WithValue":    reflect.ValueOf(q.WithValue),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
