package keyboard_shortcuts_inhibit

//go:generate go run github.com/raitonoberu/go-wayland/cmd/go-wayland-scanner -pkg keyboard_shortcuts_inhibit -prefix zwp -suffix v1 -o keyboard_shortcuts_inhibit.go -i https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/$WAYLAND_PROTOCOLS_TAG/unstable/keyboard-shortcuts-inhibit/keyboard-shortcuts-inhibit-unstable-v1.xml
