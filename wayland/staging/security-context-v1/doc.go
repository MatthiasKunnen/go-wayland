package security_context

//go:generate go run github.com/MatthiasKunnen/go-wayland/cmd/go-wayland-scanner -pkg security_context -suffix v1 -o security_context.go -i https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/1.41/staging/security-context/security-context-v1.xml
