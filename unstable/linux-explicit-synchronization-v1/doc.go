package linux_explicit_synchronization

//go:generate go run github.com/raitonoberu/go-wayland/cmd/go-wayland-scanner -pkg linux_explicit_synchronization -prefix zwp -suffix v1 -o linux_explicit_synchronization.go -i https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/$WAYLAND_PROTOCOLS_TAG/unstable/linux-explicit-synchronization/linux-explicit-synchronization-unstable-v1.xml
