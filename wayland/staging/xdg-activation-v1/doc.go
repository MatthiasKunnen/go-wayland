package xdg_activation

//go:generate go run github.com/MatthiasKunnen/go-wayland/cmd/go-wayland-scanner -pkg xdg_activation -prefix xdg -suffix v1 -o xdg_activation.go -i https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/1.41/staging/xdg-activation/xdg-activation-v1.xml?ref_type=tags
