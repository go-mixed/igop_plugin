//go:build go1.18
// +build go1.18

package main

/*

go get golang.org/x/crypto/... golang.org/x/image/... golang.org/x/net/... golang.org/x/sync/... golang.org/x/sys/... golang.org/x/text/... golang.org/x/time/...
go list golang.org/x/crypto/... golang.org/x/image/... golang.org/x/net/... golang.org/x/sync/... golang.org/x/sys/... golang.org/x/text/... golang.org/x/time/...

qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/acme
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/acme/autocert
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/acme/autocert/internal/acmetest
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/acme/internal/acmeprobe
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/argon2
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/bcrypt
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/blake2b
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/blake2s
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/blowfish
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/bn256
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/cast5
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/chacha20
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/chacha20poly1305
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/cryptobyte
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/cryptobyte/asn1
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/curve25519
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/curve25519/internal/field
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/ed25519
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/hkdf
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/internal/alias
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/internal/poly1305
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/internal/wycheproof
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/internal/wycheproof/internal/dsa
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/md4
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/nacl/auth
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/nacl/box
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/nacl/secretbox
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/nacl/sign
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/ocsp
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/openpgp
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/openpgp/armor
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/openpgp/clearsign
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/openpgp/elgamal
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/openpgp/errors
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/openpgp/packet
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/openpgp/s2k
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/otr
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/pbkdf2
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/pkcs12
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/pkcs12/internal/rc2
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/poly1305
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/ripemd160
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/salsa20
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/salsa20/salsa
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/scrypt
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/sha3
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/ssh
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/ssh/agent
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/ssh/internal/bcrypt_pbkdf
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/ssh/knownhosts
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/ssh/terminal
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/ssh/test
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/tea
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/twofish
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/xtea
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/crypto/xts
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/image/bmp
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/image/ccitt
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/image/colornames
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/image/draw
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/image/font
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/image/font/basicfont
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/image/font/gofont/gobold
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/image/font/gofont/gobolditalic
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/image/font/gofont/goitalic
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/image/font/gofont/gomedium
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/image/font/gofont/gomediumitalic
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/image/font/gofont/gomono
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/image/font/gofont/gomonobold
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/image/font/gofont/gomonobolditalic
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/image/font/gofont/gomonoitalic
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/image/font/gofont/goregular
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/image/font/gofont/gosmallcaps
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/image/font/gofont/gosmallcapsitalic
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/image/font/inconsolata
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/image/font/opentype
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/image/font/plan9font
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/image/font/sfnt
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/image/math/f32
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/image/math/f64
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/image/math/fixed
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/image/riff
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/image/tiff
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/image/tiff/lzw
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/image/vector
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/image/vp8
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/image/vp8l
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/image/webp
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/net/bpf
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/net/context
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/net/context/ctxhttp
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/net/dict
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/net/dns/dnsmessage
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/net/html
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/net/html/atom
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/net/html/charset
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/net/http/httpguts
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/net/http/httpproxy
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/net/http2
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/net/http2/h2c
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/net/http2/h2i
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/net/http2/hpack
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/net/icmp
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/net/idna
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/net/internal/iana
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/net/internal/socket
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/net/internal/socks
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/net/internal/sockstest
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/net/internal/timeseries
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/net/ipv4
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/net/ipv6
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/net/nettest
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/net/netutil
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/net/proxy
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/net/publicsuffix
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/net/trace
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/net/webdav
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/net/webdav/internal/xml
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/net/websocket
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/net/xsrftoken
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/sync/errgroup
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/sync/semaphore
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/sync/singleflight
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/sync/syncmap
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/sys/cpu
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/sys/execabs
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/sys/internal/unsafeheader
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/sys/unix
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/sys/unix/internal/mkmerge
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/sys/windows/mkwinsyscall
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/cases
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/cmd/gotext
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/cmd/gotext/examples/extract
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/cmd/gotext/examples/extract_http
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/cmd/gotext/examples/extract_http/pkg
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/cmd/gotext/examples/rewrite
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/collate
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/collate/build
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/collate/tools/colcmp
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/currency
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/date
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/encoding
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/encoding/charmap
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/encoding/htmlindex
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/encoding/ianaindex
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/encoding/internal
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/encoding/internal/enctest
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/encoding/internal/identifier
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/encoding/japanese
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/encoding/korean
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/encoding/simplifiedchinese
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/encoding/traditionalchinese
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/encoding/unicode
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/encoding/unicode/utf32
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/feature/plural
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/internal
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/internal/catmsg
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/internal/cldrtree
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/internal/colltab
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/internal/export/idna
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/internal/export/unicode
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/internal/format
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/internal/gen
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/internal/gen/bitfield
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/internal/language
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/internal/language/compact
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/internal/number
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/internal/stringset
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/internal/tag
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/internal/testtext
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/internal/triegen
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/internal/ucd
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/internal/utf8internal
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/language
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/language/display
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/message
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/message/catalog
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/message/pipeline
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/number
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/runes
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/search
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/secure
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/secure/bidirule
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/secure/precis
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/transform
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/unicode
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/unicode/bidi
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/unicode/cldr
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/unicode/norm
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/unicode/rangetable
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/unicode/runenames
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/text/width
qexp -outdir . -addtags "//+build go1.18" -filename go118_export golang.org/x/time/rate
*/

import (
	_ "igop_plugin_x/golang.org/x/crypto/acme"
	_ "igop_plugin_x/golang.org/x/crypto/acme/autocert"
	_ "igop_plugin_x/golang.org/x/crypto/argon2"
	_ "igop_plugin_x/golang.org/x/crypto/bcrypt"
	_ "igop_plugin_x/golang.org/x/crypto/blake2b"
	_ "igop_plugin_x/golang.org/x/crypto/blake2s"
	_ "igop_plugin_x/golang.org/x/crypto/blowfish"
	_ "igop_plugin_x/golang.org/x/crypto/bn256"
	_ "igop_plugin_x/golang.org/x/crypto/cast5"
	_ "igop_plugin_x/golang.org/x/crypto/chacha20"
	_ "igop_plugin_x/golang.org/x/crypto/chacha20poly1305"
	_ "igop_plugin_x/golang.org/x/crypto/cryptobyte"
	_ "igop_plugin_x/golang.org/x/crypto/cryptobyte/asn1"
	_ "igop_plugin_x/golang.org/x/crypto/curve25519"
	_ "igop_plugin_x/golang.org/x/crypto/ed25519"
	_ "igop_plugin_x/golang.org/x/crypto/hkdf"
	_ "igop_plugin_x/golang.org/x/crypto/md4"
	_ "igop_plugin_x/golang.org/x/crypto/nacl/auth"
	_ "igop_plugin_x/golang.org/x/crypto/nacl/box"
	_ "igop_plugin_x/golang.org/x/crypto/nacl/secretbox"
	_ "igop_plugin_x/golang.org/x/crypto/nacl/sign"
	_ "igop_plugin_x/golang.org/x/crypto/ocsp"
	_ "igop_plugin_x/golang.org/x/crypto/openpgp"
	_ "igop_plugin_x/golang.org/x/crypto/openpgp/armor"
	_ "igop_plugin_x/golang.org/x/crypto/openpgp/clearsign"
	_ "igop_plugin_x/golang.org/x/crypto/openpgp/elgamal"
	_ "igop_plugin_x/golang.org/x/crypto/openpgp/errors"
	_ "igop_plugin_x/golang.org/x/crypto/openpgp/packet"
	_ "igop_plugin_x/golang.org/x/crypto/openpgp/s2k"
	_ "igop_plugin_x/golang.org/x/crypto/otr"
	_ "igop_plugin_x/golang.org/x/crypto/pbkdf2"
	_ "igop_plugin_x/golang.org/x/crypto/pkcs12"
	_ "igop_plugin_x/golang.org/x/crypto/poly1305"
	_ "igop_plugin_x/golang.org/x/crypto/ripemd160"
	_ "igop_plugin_x/golang.org/x/crypto/salsa20"
	_ "igop_plugin_x/golang.org/x/crypto/salsa20/salsa"
	_ "igop_plugin_x/golang.org/x/crypto/scrypt"
	_ "igop_plugin_x/golang.org/x/crypto/sha3"
	_ "igop_plugin_x/golang.org/x/crypto/ssh"
	_ "igop_plugin_x/golang.org/x/crypto/ssh/agent"
	_ "igop_plugin_x/golang.org/x/crypto/ssh/knownhosts"
	_ "igop_plugin_x/golang.org/x/crypto/ssh/terminal"
	_ "igop_plugin_x/golang.org/x/crypto/tea"
	_ "igop_plugin_x/golang.org/x/crypto/twofish"
	_ "igop_plugin_x/golang.org/x/crypto/xtea"
	_ "igop_plugin_x/golang.org/x/crypto/xts"
	_ "igop_plugin_x/golang.org/x/image/bmp"
	_ "igop_plugin_x/golang.org/x/image/ccitt"
	_ "igop_plugin_x/golang.org/x/image/colornames"
	_ "igop_plugin_x/golang.org/x/image/draw"
	_ "igop_plugin_x/golang.org/x/image/font"
	_ "igop_plugin_x/golang.org/x/image/font/basicfont"
	_ "igop_plugin_x/golang.org/x/image/font/gofont/gobold"
	_ "igop_plugin_x/golang.org/x/image/font/gofont/gobolditalic"
	_ "igop_plugin_x/golang.org/x/image/font/gofont/goitalic"
	_ "igop_plugin_x/golang.org/x/image/font/gofont/gomedium"
	_ "igop_plugin_x/golang.org/x/image/font/gofont/gomediumitalic"
	_ "igop_plugin_x/golang.org/x/image/font/gofont/gomono"
	_ "igop_plugin_x/golang.org/x/image/font/gofont/gomonobold"
	_ "igop_plugin_x/golang.org/x/image/font/gofont/gomonobolditalic"
	_ "igop_plugin_x/golang.org/x/image/font/gofont/gomonoitalic"
	_ "igop_plugin_x/golang.org/x/image/font/gofont/goregular"
	_ "igop_plugin_x/golang.org/x/image/font/gofont/gosmallcaps"
	_ "igop_plugin_x/golang.org/x/image/font/gofont/gosmallcapsitalic"
	_ "igop_plugin_x/golang.org/x/image/font/inconsolata"
	_ "igop_plugin_x/golang.org/x/image/font/opentype"
	_ "igop_plugin_x/golang.org/x/image/font/plan9font"
	_ "igop_plugin_x/golang.org/x/image/font/sfnt"
	_ "igop_plugin_x/golang.org/x/image/math/f32"
	_ "igop_plugin_x/golang.org/x/image/math/f64"
	_ "igop_plugin_x/golang.org/x/image/math/fixed"
	_ "igop_plugin_x/golang.org/x/image/riff"
	_ "igop_plugin_x/golang.org/x/image/tiff"
	_ "igop_plugin_x/golang.org/x/image/tiff/lzw"
	_ "igop_plugin_x/golang.org/x/image/vector"
	_ "igop_plugin_x/golang.org/x/image/vp8"
	_ "igop_plugin_x/golang.org/x/image/vp8l"
	_ "igop_plugin_x/golang.org/x/image/webp"
	_ "igop_plugin_x/golang.org/x/net/bpf"
	_ "igop_plugin_x/golang.org/x/net/context"
	_ "igop_plugin_x/golang.org/x/net/context/ctxhttp"
	_ "igop_plugin_x/golang.org/x/net/dict"
	_ "igop_plugin_x/golang.org/x/net/dns/dnsmessage"
	_ "igop_plugin_x/golang.org/x/net/html"
	_ "igop_plugin_x/golang.org/x/net/html/atom"
	_ "igop_plugin_x/golang.org/x/net/html/charset"
	_ "igop_plugin_x/golang.org/x/net/http/httpguts"
	_ "igop_plugin_x/golang.org/x/net/http/httpproxy"
	_ "igop_plugin_x/golang.org/x/net/http2"
	_ "igop_plugin_x/golang.org/x/net/http2/h2c"
	_ "igop_plugin_x/golang.org/x/net/http2/hpack"
	_ "igop_plugin_x/golang.org/x/net/icmp"
	_ "igop_plugin_x/golang.org/x/net/idna"
	_ "igop_plugin_x/golang.org/x/net/ipv4"
	_ "igop_plugin_x/golang.org/x/net/ipv6"
	_ "igop_plugin_x/golang.org/x/net/nettest"
	_ "igop_plugin_x/golang.org/x/net/netutil"
	_ "igop_plugin_x/golang.org/x/net/proxy"
	_ "igop_plugin_x/golang.org/x/net/publicsuffix"
	_ "igop_plugin_x/golang.org/x/net/trace"
	_ "igop_plugin_x/golang.org/x/net/webdav"
	_ "igop_plugin_x/golang.org/x/net/websocket"
	_ "igop_plugin_x/golang.org/x/net/xsrftoken"
	_ "igop_plugin_x/golang.org/x/sync/errgroup"
	_ "igop_plugin_x/golang.org/x/sync/semaphore"
	_ "igop_plugin_x/golang.org/x/sync/singleflight"
	_ "igop_plugin_x/golang.org/x/sync/syncmap"
	_ "igop_plugin_x/golang.org/x/sys/cpu"
	_ "igop_plugin_x/golang.org/x/sys/execabs"
	_ "igop_plugin_x/golang.org/x/text/cases"
	_ "igop_plugin_x/golang.org/x/text/collate"
	_ "igop_plugin_x/golang.org/x/text/collate/build"
	_ "igop_plugin_x/golang.org/x/text/currency"
	_ "igop_plugin_x/golang.org/x/text/date"
	_ "igop_plugin_x/golang.org/x/text/encoding"
	_ "igop_plugin_x/golang.org/x/text/encoding/charmap"
	_ "igop_plugin_x/golang.org/x/text/encoding/htmlindex"
	_ "igop_plugin_x/golang.org/x/text/encoding/ianaindex"
	_ "igop_plugin_x/golang.org/x/text/encoding/japanese"
	_ "igop_plugin_x/golang.org/x/text/encoding/korean"
	_ "igop_plugin_x/golang.org/x/text/encoding/simplifiedchinese"
	_ "igop_plugin_x/golang.org/x/text/encoding/traditionalchinese"
	_ "igop_plugin_x/golang.org/x/text/encoding/unicode"
	_ "igop_plugin_x/golang.org/x/text/encoding/unicode/utf32"
	_ "igop_plugin_x/golang.org/x/text/feature/plural"
	_ "igop_plugin_x/golang.org/x/text/language"
	_ "igop_plugin_x/golang.org/x/text/language/display"
	_ "igop_plugin_x/golang.org/x/text/message"
	_ "igop_plugin_x/golang.org/x/text/message/catalog"
	_ "igop_plugin_x/golang.org/x/text/message/pipeline"
	_ "igop_plugin_x/golang.org/x/text/number"
	_ "igop_plugin_x/golang.org/x/text/runes"
	_ "igop_plugin_x/golang.org/x/text/search"
	_ "igop_plugin_x/golang.org/x/text/secure/bidirule"
	_ "igop_plugin_x/golang.org/x/text/secure/precis"
	_ "igop_plugin_x/golang.org/x/text/transform"
	_ "igop_plugin_x/golang.org/x/text/unicode/bidi"
	_ "igop_plugin_x/golang.org/x/text/unicode/cldr"
	_ "igop_plugin_x/golang.org/x/text/unicode/norm"
	_ "igop_plugin_x/golang.org/x/text/unicode/rangetable"
	_ "igop_plugin_x/golang.org/x/text/unicode/runenames"
	_ "igop_plugin_x/golang.org/x/text/width"
	_ "igop_plugin_x/golang.org/x/time/rate"
)

func Load() {

}