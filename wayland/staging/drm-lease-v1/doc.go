package drm_lease

//go:generate go run github.com/MatthiasKunnen/go-wayland/cmd/go-wayland-scanner -pkg drm_lease -prefix wp -suffix v1 -o drm_lease.go -i https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/1.45/staging/drm-lease/drm-lease-v1.xml?ref_type=tags
