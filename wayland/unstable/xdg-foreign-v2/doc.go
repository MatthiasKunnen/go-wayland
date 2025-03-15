package xdg_foreign

//go:generate go run github.com/MatthiasKunnen/go-wayland/cmd/go-wayland-scanner -pkg xdg_foreign -prefix zxdg -suffix v2 -o xdg_foreign.go -i https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/1.41/unstable/xdg-foreign/xdg-foreign-unstable-v2.xml?ref_type=tags
