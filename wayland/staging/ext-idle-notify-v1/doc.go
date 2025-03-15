package ext_idle_notify

//go:generate go run github.com/MatthiasKunnen/go-wayland/cmd/go-wayland-scanner -pkg ext_idle_notify -prefix ext -suffix v1 -o ext_idle_notify.go -i https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/1.41/staging/ext-idle-notify/ext-idle-notify-v1.xml?ref_type=tags
