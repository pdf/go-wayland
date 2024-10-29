package fullscreen_shell

//go:generate go run github.com/raitonoberu/go-wayland/cmd/go-wayland-scanner -pkg fullscreen_shell -prefix zwp -suffix v1 -o fullscreen_shell.go -i https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/$WAYLAND_PROTOCOLS_TAG/unstable/fullscreen-shell/fullscreen-shell-unstable-v1.xml
