package text_input

//go:generate go run github.com/pdf/go-wayland/cmd/go-wayland-scanner -pkg text_input -prefix zwp -suffix v3 -o text_input.go -i https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/$WAYLAND_PROTOCOLS_TAG/unstable/text-input/text-input-unstable-v3.xml
