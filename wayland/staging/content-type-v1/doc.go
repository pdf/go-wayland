package content_type

//go:generate go run github.com/rajveermalviya/go-wayland/cmd/go-wayland-scanner -pkg content_type -prefix wp -suffix v1 -o content_type.go -i https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/$WAYLAND_PROTOCOLS_TAG/staging/content-type/content-type-v1.xml
