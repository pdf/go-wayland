package input_method

//go:generate go run github.com/rajveermalviya/go-wayland/cmd/go-wayland-scanner -pkg input_method -prefix zwp -suffix v1 -o input_method.go -i https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/$WAYLAND_PROTOCOLS_TAG/unstable/input-method/input-method-unstable-v1.xml
