package pointer_constraints

//go:generate go run github.com/pdf/go-wayland/cmd/go-wayland-scanner -pkg pointer_constraints -prefix zwp -suffix v1 -o pointer_constraints.go -i https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/$WAYLAND_PROTOCOLS_TAG/unstable/pointer-constraints/pointer-constraints-unstable-v1.xml
