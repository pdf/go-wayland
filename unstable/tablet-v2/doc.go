package tablet

//go:generate go run github.com/pdf/go-wayland/cmd/go-wayland-scanner -pkg tablet -prefix zwp -suffix v2 -o tablet.go -i https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/$WAYLAND_PROTOCOLS_TAG/unstable/tablet/tablet-unstable-v2.xml
