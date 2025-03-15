package content_type

//go:generate go run github.com/MatthiasKunnen/go-wayland/cmd/go-wayland-scanner -pkg content_type -prefix wp -suffix v1 -o content_type.go -i https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/1.41/staging/content-type/content-type-v1.xml?ref_type=tags
