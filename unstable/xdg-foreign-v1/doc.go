package xdg_foreign

//go:generate go run github.com/pdf/go-wayland/cmd/go-wayland-scanner -pkg xdg_foreign -prefix zxdg -suffix v1 -o xdg_foreign.go -i https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/$WAYLAND_PROTOCOLS_TAG/unstable/xdg-foreign/xdg-foreign-unstable-v1.xml
