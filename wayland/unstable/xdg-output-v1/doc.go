package xdg_output

//go:generate go run github.com/MatthiasKunnen/go-wayland/cmd/go-wayland-scanner -pkg xdg_output -prefix zxdg -suffix v1 -o xdg_output.go -i https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/1.41/unstable/xdg-output/xdg-output-unstable-v1.xml?ref_type=tags
