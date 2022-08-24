# igop plugins

Use with https://github.com/fly-studio/igop, built to plugin to reduce the igop file size.

When assembly files "*.s" are present in a 3<sup>rd</sup> party package, igop can not dynamically import them, only if be compiled to igop, but it will inflate the igop file size.

So compilation into plug-in mode can be imported on demand.

> Only for Linux, Mac, BSD.

## Build

```
$ cd x/
$ go build -buildmode=plugin
```

It'll building `igop_plugin_x.so`

## Usage

```
$ ./igop run -p /path/to/igop_plugin_x.so /path/to/example4
```

## Folders

- **x** : golang.org/x/...
