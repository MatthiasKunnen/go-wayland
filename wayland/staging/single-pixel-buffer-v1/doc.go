package single_pixel_buffer

//go:generate go run github.com/MatthiasKunnen/go-wayland/cmd/go-wayland-scanner -pkg single_pixel_buffer -suffix v1 -o single_pixel_buffer.go -i https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/1.41/staging/single-pixel-buffer/single-pixel-buffer-v1.xml?ref_type=tags
