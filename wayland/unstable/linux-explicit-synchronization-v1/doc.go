package linux_explicit_synchronization

//go:generate go run github.com/MatthiasKunnen/go-wayland/cmd/go-wayland-scanner -pkg linux_explicit_synchronization -prefix zwp -suffix v1 -o linux_explicit_synchronization.go -i https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/1.41/unstable/linux-explicit-synchronization/linux-explicit-synchronization-unstable-v1.xml?ref_type=tags
