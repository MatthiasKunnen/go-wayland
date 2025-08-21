package pointer_constraints

//go:generate go run github.com/MatthiasKunnen/go-wayland/cmd/go-wayland-scanner -pkg pointer_constraints -prefix zwp -suffix v1 -o pointer_constraints.go -i https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/1.45/unstable/pointer-constraints/pointer-constraints-unstable-v1.xml?ref_type=tags
