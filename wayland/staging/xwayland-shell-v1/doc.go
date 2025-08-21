package xwayland_shell

//go:generate go run github.com/MatthiasKunnen/go-wayland/cmd/go-wayland-scanner -pkg xwayland_shell -suffix v1 -o xwayland_shell.go -i https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/1.45/staging/xwayland-shell/xwayland-shell-v1.xml?ref_type=tags
