# Wayland implementation in Go

[![Go Reference](https://pkg.go.dev/badge/github.com/MatthiasKunnen/go-wayland/wayland.svg)](https://pkg.go.dev/github.com/MatthiasKunnen/go-wayland/wayland)

This module contains pure Go implementation of the Wayland protocol.
Currently only wayland-client functionality is supported.

Go code is generated from protocol XML files using
[`go-wayland-scanner`](cmd/go-wayland-scanner/scanner.go).

To load cursor, minimal port of `wayland-cursor` & `xcursor` in pure Go
is located at [`wayland/cursor`](wayland/cursor) & [`wayland/cursor/xcursor`](wayland/cursor/xcursor)
respectively.

To demonstrate the functionality of this module
[`examples/imageviewer`](examples/imageviewer) contains a simple image
viewer. It demos displaying a top-level window, resizing of window,
cursor themes, pointer & keyboard. Because it's in pure Go, it can be
compiled without CGO. You can try it using following commands:

```sh
CGO_ENABLED=0 go install github.com/MatthiasKunnen/go-wayland/examples/imageviewer@latest

imageviewer file.jpg
```

## Fork
This repository is forked from [github.com/rajveermalviya/go-wayland](https://github.com/rajveermalviya/go-wayland) and includes several fixes and improvements.
Forked at `0ad78d1`, see [the diff](https://github.com/MatthiasKunnen/go-wayland/compare/0ad78d1...main).
