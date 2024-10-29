package tablet

//go:generate go run github.com/raitonoberu/go-wayland/cmd/go-wayland-scanner -pkg tablet -prefix zwp -suffix v1 -o tablet.go -i https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/$WAYLAND_PROTOCOLS_TAG/unstable/tablet/tablet-unstable-v1.xml
