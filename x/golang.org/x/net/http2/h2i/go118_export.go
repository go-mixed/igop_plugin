// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18

package main

import (
	_ "golang.org/x/net/http2/h2i"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "main",
		Path: "golang.org/x/net/http2/h2i",
		Deps: map[string]string{
			"bufio":                        "bufio",
			"bytes":                        "bytes",
			"crypto/tls":                   "tls",
			"errors":                       "errors",
			"flag":                         "flag",
			"fmt":                          "fmt",
			"golang.org/x/net/http2":       "http2",
			"golang.org/x/net/http2/hpack": "hpack",
			"golang.org/x/term":            "term",
			"io":                           "io",
			"log":                          "log",
			"net":                          "net",
			"net/http":                     "http",
			"os":                           "os",
			"regexp":                       "regexp",
			"strconv":                      "strconv",
			"strings":                      "strings",
		},
	})
}
