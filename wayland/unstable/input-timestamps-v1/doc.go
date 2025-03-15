package input_timestamps

//go:generate go run github.com/MatthiasKunnen/go-wayland/cmd/go-wayland-scanner -pkg input_timestamps -prefix zwp -suffix v1 -o input_timestamps.go -i https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/1.41/unstable/input-timestamps/input-timestamps-unstable-v1.xml?ref_type=tags
