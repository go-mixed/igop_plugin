package main

/*

go get golang.org/x/crypto/... golang.org/x/image/... golang.org/x/net/... golang.org/x/sync/... golang.org/x/sys/... golang.org/x/text/... golang.org/x/time/...
go list golang.org/x/crypto/... golang.org/x/image/... golang.org/x/net/... golang.org/x/sync/... golang.org/x/sys/... golang.org/x/text/... golang.org/x/time/...

set GOOS=linux
qexp -outdir . -addtags "//go:build go1.19" -filename go119_export golang.org/x/crypto/acme golang.org/x/crypto/acme/autocert golang.org/x/crypto/argon2 golang.org/x/crypto/bcrypt golang.org/x/crypto/blake2b golang.org/x/crypto/blake2s golang.org/x/crypto/blowfish golang.org/x/crypto/bn256 golang.org/x/crypto/cast5 golang.org/x/crypto/chacha20 golang.org/x/crypto/chacha20poly1305 golang.org/x/crypto/cryptobyte golang.org/x/crypto/cryptobyte/asn1 golang.org/x/crypto/curve25519 golang.org/x/crypto/ed25519 golang.org/x/crypto/hkdf golang.org/x/crypto/md4 golang.org/x/crypto/nacl/auth golang.org/x/crypto/nacl/box golang.org/x/crypto/nacl/secretbox golang.org/x/crypto/nacl/sign golang.org/x/crypto/ocsp golang.org/x/crypto/openpgp golang.org/x/crypto/openpgp/armor golang.org/x/crypto/openpgp/clearsign golang.org/x/crypto/openpgp/elgamal golang.org/x/crypto/openpgp/errors golang.org/x/crypto/openpgp/packet golang.org/x/crypto/openpgp/s2k golang.org/x/crypto/otr golang.org/x/crypto/pbkdf2 golang.org/x/crypto/pkcs12 golang.org/x/crypto/poly1305 golang.org/x/crypto/ripemd160 golang.org/x/crypto/salsa20 golang.org/x/crypto/salsa20/salsa golang.org/x/crypto/scrypt golang.org/x/crypto/sha3 golang.org/x/crypto/ssh golang.org/x/crypto/ssh/agent golang.org/x/crypto/ssh/knownhosts golang.org/x/crypto/ssh/terminal golang.org/x/crypto/ssh/test golang.org/x/crypto/tea golang.org/x/crypto/twofish golang.org/x/crypto/xtea golang.org/x/crypto/xts golang.org/x/image/bmp golang.org/x/image/ccitt golang.org/x/image/colornames golang.org/x/image/draw golang.org/x/image/font golang.org/x/image/font/basicfont golang.org/x/image/font/gofont/gobold golang.org/x/image/font/gofont/gobolditalic golang.org/x/image/font/gofont/goitalic golang.org/x/image/font/gofont/gomedium golang.org/x/image/font/gofont/gomediumitalic golang.org/x/image/font/gofont/gomono golang.org/x/image/font/gofont/gomonobold golang.org/x/image/font/gofont/gomonobolditalic golang.org/x/image/font/gofont/gomonoitalic golang.org/x/image/font/gofont/goregular golang.org/x/image/font/gofont/gosmallcaps golang.org/x/image/font/gofont/gosmallcapsitalic golang.org/x/image/font/inconsolata golang.org/x/image/font/opentype golang.org/x/image/font/plan9font golang.org/x/image/font/sfnt golang.org/x/image/math/f32 golang.org/x/image/math/f64 golang.org/x/image/math/fixed golang.org/x/image/riff golang.org/x/image/tiff golang.org/x/image/tiff/lzw golang.org/x/image/vector golang.org/x/image/vp8 golang.org/x/image/vp8l golang.org/x/image/webp golang.org/x/net/bpf golang.org/x/net/context golang.org/x/net/context/ctxhttp golang.org/x/net/dict golang.org/x/net/dns/dnsmessage golang.org/x/net/html golang.org/x/net/html/atom golang.org/x/net/html/charset golang.org/x/net/http/httpguts golang.org/x/net/http/httpproxy golang.org/x/net/http2 golang.org/x/net/http2/h2c golang.org/x/net/http2/h2i golang.org/x/net/http2/hpack golang.org/x/net/icmp golang.org/x/net/idna golang.org/x/net/ipv4 golang.org/x/net/ipv6 golang.org/x/net/nettest golang.org/x/net/netutil golang.org/x/net/proxy golang.org/x/net/publicsuffix golang.org/x/net/trace golang.org/x/net/webdav golang.org/x/net/websocket golang.org/x/net/xsrftoken golang.org/x/sync/errgroup golang.org/x/sync/semaphore golang.org/x/sync/singleflight golang.org/x/sync/syncmap golang.org/x/sys/cpu golang.org/x/sys/execabs golang.org/x/sys/unix golang.org/x/text golang.org/x/text/cases golang.org/x/text/cmd/gotext golang.org/x/text/cmd/gotext/examples/extract golang.org/x/text/cmd/gotext/examples/extract_http golang.org/x/text/cmd/gotext/examples/extract_http/pkg golang.org/x/text/cmd/gotext/examples/rewrite golang.org/x/text/collate golang.org/x/text/collate/build golang.org/x/text/collate/tools/colcmp golang.org/x/text/currency golang.org/x/text/date golang.org/x/text/encoding golang.org/x/text/encoding/charmap golang.org/x/text/encoding/htmlindex golang.org/x/text/encoding/ianaindex golang.org/x/text/encoding/japanese golang.org/x/text/encoding/korean golang.org/x/text/encoding/simplifiedchinese golang.org/x/text/encoding/traditionalchinese golang.org/x/text/encoding/unicode golang.org/x/text/encoding/unicode/utf32 golang.org/x/text/feature/plural golang.org/x/text/language golang.org/x/text/language/display golang.org/x/text/message golang.org/x/text/message/catalog golang.org/x/text/message/pipeline golang.org/x/text/number golang.org/x/text/runes golang.org/x/text/search golang.org/x/text/secure golang.org/x/text/secure/bidirule golang.org/x/text/secure/precis golang.org/x/text/transform golang.org/x/text/unicode golang.org/x/text/unicode/bidi golang.org/x/text/unicode/cldr golang.org/x/text/unicode/norm golang.org/x/text/unicode/rangetable golang.org/x/text/unicode/runenames golang.org/x/text/width golang.org/x/time/rate

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
