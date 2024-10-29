package viewporter

//go:generate go run github.com/raitonoberu/go-wayland/cmd/go-wayland-scanner -pkg viewporter -prefix wp -o viewporter.go -i https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/$WAYLAND_PROTOCOLS_TAG/stable/viewporter/viewporter.xml
