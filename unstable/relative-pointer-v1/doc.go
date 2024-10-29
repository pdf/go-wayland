package relative_pointer

//go:generate go run github.com/raitonoberu/go-wayland/cmd/go-wayland-scanner -pkg relative_pointer -prefix zwp -suffix v1 -o relative_pointer.go -i https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/$WAYLAND_PROTOCOLS_TAG/unstable/relative-pointer/relative-pointer-unstable-v1.xml
