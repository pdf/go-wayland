package primary_selection

//go:generate go run github.com/raitonoberu/go-wayland/cmd/go-wayland-scanner -pkg primary_selection -prefix zwp -suffix v1 -o primary_selection.go -i https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/$WAYLAND_PROTOCOLS_TAG/unstable/primary-selection/primary-selection-unstable-v1.xml
