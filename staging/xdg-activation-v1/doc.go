package xdg_activation

//go:generate go run github.com/raitonoberu/go-wayland/cmd/go-wayland-scanner -pkg xdg_activation -prefix xdg -suffix v1 -o xdg_activation.go -i https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/$WAYLAND_PROTOCOLS_TAG/staging/xdg-activation/xdg-activation-v1.xml
