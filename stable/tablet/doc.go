package tablet

//go:generate go run github.com/raitonoberu/go-wayland/cmd/go-wayland-scanner -pkg tablet -prefix zwp -suffix v2 -o tablet.go -i https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/$WAYLAND_PROTOCOLS_TAG/stable/tablet/tablet-v2.xml
