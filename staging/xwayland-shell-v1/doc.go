package xwayland_shell

//go:generate go run github.com/raitonoberu/go-wayland/cmd/go-wayland-scanner -pkg xwayland_shell -suffix v1 -o xwayland_shell.go -i https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/$WAYLAND_PROTOCOLS_TAG/staging/xwayland-shell/xwayland-shell-v1.xml
