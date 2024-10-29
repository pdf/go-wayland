package xdg_shell

//go:generate go run github.com/rajveermalviya/go-wayland/cmd/go-wayland-scanner -pkg xdg_shell -prefix xdg  -o xdg_shell.go -i https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/$WAYLAND_PROTOCOLS_TAG/stable/xdg-shell/xdg-shell.xml
